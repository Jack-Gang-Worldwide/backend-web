package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/valorant/entity"
	sq "github.com/Masterminds/squirrel"
)

type valorantRepositoryImpl struct {
	DB *sql.DB
}

func ProvideValorantRepository(DB *sql.DB) *valorantRepositoryImpl{
	return &valorantRepositoryImpl{DB: DB}
}

var(
	SELECT_VALORANT_MEMBERS = sq.Select("valorant.id", "valorant.member_id", "valorant.ingame_name", "member.real_name", "valorant.image", "valorant.role", "valorant.nationality", "valorant.created_at", "valorant.updated_at").From("valorant").Join("member ON member.id = valorant.member_id")
	SELECT_REAL_NAME =sq.Select("real_name").From("member")
	CHECK_VALORANT_MEMBERS = sq.Select("count(*)").From("valorant")
)

func(vr valorantRepositoryImpl)GetAllMembers(ctx context.Context)(entity.ValorantMembers, entity.Members, error){
	query := SELECT_VALORANT_MEMBERS.OrderBy("created_at ASC")
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}
	prpd, err := vr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}
	rows, err := prpd.Query()
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}
	membersValorant := entity.ValorantMembers{}
	members := entity.Members{}
	for rows.Next(){
		var memberValorant entity.ValorantMember
		var member entity.Member

		err := rows.Scan(&memberValorant.Id, &memberValorant.MemberId, &memberValorant.InGameName, &member.RealName, &memberValorant.Image, &memberValorant.Role, &memberValorant.Nationality, &memberValorant.CreatedAt, &memberValorant.UpdatedAt)
		if err != nil {
			log.Printf("ERROR GetAllMembers -> error: %v\n", err)
			return nil, nil, err
		}
		membersValorant = append(membersValorant, &memberValorant)
		members = append(members, &member)
	}
	return membersValorant, members, nil
}

func(vr valorantRepositoryImpl)CheckMembersExists(ctx context.Context)(int, error){
	query := CHECK_VALORANT_MEMBERS
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR CheckMembersExists -> error: %v\n", err)
		return 0, err
	}
	prpd, err := vr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("ERROR CheckMembersExists -> error: %v\n", err)
		return 0, err
	}
	rows, err := prpd.Query()
	if err != nil {
		log.Printf("ERROR CheckMembersExists -> error: %v\n", err)
		return 0, err
	}
	var membersCount int
	if rows.Next(){
		err := rows.Scan(&membersCount)
		if err != nil {
			log.Printf("ERROR CheckMembersExists -> error: %v\n", err)
			return 0, err
		}
	}
	return membersCount, nil
}

func(vr valorantRepositoryImpl)GetListMembers(ctx context.Context)(entity.Members, error){
	query := SELECT_REAL_NAME.OrderBy("created_at ASC")
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, err
	}
	prpd, err := vr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, err
	}
	rows, err := prpd.Query()
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, err
	}

	members := entity.Members{}

	for rows.Next(){
		var member entity.Member

		err := rows.Scan(&member.Id, &member.RealName, &member.Birth, &member.Born, &member.Description, &member.CreatedAt, &member.UpdatedAt)
		if err != nil {
			log.Printf("ERROR GetAllMembers -> error: %v\n", err)
			return nil, err
		}

		members = append(members, &member)
	}
	return members, nil
}