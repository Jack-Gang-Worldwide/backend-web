package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/apexlegends/entity"
	sq "github.com/Masterminds/squirrel"
)

type apexLegendsRepositoryImpl struct {
	DB *sql.DB
}

func ProvideApexLegendsRepository(DB *sql.DB) *apexLegendsRepositoryImpl{
	return &apexLegendsRepositoryImpl{DB: DB}
}

var(
	SELECT_APEX_MEMBERS = sq.Select("apex_legends.id", "apex_legends.member_id", "apex_legends.ingame_name", "member.real_name", "apex_legends.image", "apex_legends.nationality", "apex_legends.created_at", "apex_legends.updated_at").From("apex_legends").Join("member ON member.id = apex_legends.id")
	SELECT_REAL_NAME = sq.Select("real_name").From("member")
	CHECK_APEX_EXISTS = sq.Select("count(*)").From("apex_legends")
)

func(ar apexLegendsRepositoryImpl)GetAllMembers(ctx context.Context)(entity.ApexLegendsMembers, entity.Members, error){
	query := SELECT_APEX_MEMBERS.OrderBy("created_at")
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}
	prpd, err := ar.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}
	rows, err := prpd.Query()
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}
	membersApex := entity.ApexLegendsMembers{}
	members := entity.Members{}
	for rows.Next(){
		var memberApex entity.ApexLegendsMember
		var member entity.Member

		err := rows.Scan(&memberApex.Id, &memberApex.MemberId, &memberApex.InGameName, &member.RealName, &memberApex.Image, &memberApex.Nationality, &memberApex.CreatedAt, &memberApex.UpdatedAt)
		if err != nil {
			log.Printf("ERROR GetAllMembers -> error: %v\n", err)
			return nil, nil, err
		}
		membersApex = append(membersApex, &memberApex)
		members = append(members, &member)
	}
	return membersApex, members, nil
}

func(ar apexLegendsRepositoryImpl)CheckMembersExists(ctx context.Context)(int, error){
	query := CHECK_APEX_EXISTS
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR CheckMembersExists -> error: %v\n", err)
		return 0, err
	}
	prpd, err := ar.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("ERROR CheckMembersExists -> error: %v\n", err)
		return 0, err
	}
	rows, err := prpd.Query()
	if err != nil {
		log.Printf("ERROR CheckMembersExists -> error: %v\n", err)
		return 0, err
	}
	var memberCount int
	if rows.Next(){
		err := rows.Scan(&memberCount)
		if err != nil {
			log.Printf("ERROR CheckMembersExists -> error: %v\n", err)
			return 0, err
		}
	}
	return memberCount, nil
}
