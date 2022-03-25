package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/mobilelegends/entity"
	sq "github.com/Masterminds/squirrel"
)

type mobilelegendsRepositoryImpl struct {
	DB *sql.DB
}

func ProvideMobileLegendsRepository(DB *sql.DB) *mobilelegendsRepositoryImpl{
	return &mobilelegendsRepositoryImpl{DB: DB}
}

var (
	SELECT_MOBILE_LEGENDS_MEMBERS = sq.Select("mobile_legends.id", "mobile_legends.member_id", "mobile_legends.ingame_name", "member.real_name", "mobile_legends.image", "mobile_legends.role", "mobile_legends.nationality", "mobile_legends.created_at", "mobile_legends.updated_at").From("mobile_legends").Join("member ON member.id = mobile_legends.member_id")
	SELECT_REAL_NAME =sq.Select("real_name").From("member")
	CHECK_MOBILE_LEGENDS_MEMBERS = sq.Select("count(*)").From("mobile_legends")
)

func(mr mobilelegendsRepositoryImpl)GetAllMembers(ctx context.Context)(entity.MobileLegendsMembers, entity.Members, error){
	query := SELECT_MOBILE_LEGENDS_MEMBERS.OrderBy("created_at ASC")
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}
	prpd, err := mr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}
	rows, err := prpd.Query()
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}

	membersMobileLegends := entity.MobileLegendsMembers{}
	members := entity.Members{}

	for rows.Next(){
		var memberMobileLegends entity.MobileLegendsMember
		var member entity.Member

		err := rows.Scan(&memberMobileLegends.Id, &memberMobileLegends.MemberId, &memberMobileLegends.InGameName, &member.RealName, &memberMobileLegends.Image, &memberMobileLegends.Role, &memberMobileLegends.Nationality, &memberMobileLegends.CreatedAt, &memberMobileLegends.UpdatedAt)
		if err != nil {
			log.Printf("ERROR GetAllMembers -> error: %v\n", err)
			return nil, nil, err
		}
		membersMobileLegends = append(membersMobileLegends, &memberMobileLegends)
		members = append(members, &member)
	}
	return membersMobileLegends, members, nil
}

func(mr mobilelegendsRepositoryImpl)CheckMembersExists(ctx context.Context)(int, error){
	query := CHECK_MOBILE_LEGENDS_MEMBERS
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR CheckMembersExists -> error: %v\n", err)
		return 0, err
	}
	prpd, err := mr.DB.PrepareContext(ctx, stmt)
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

func(mr mobilelegendsRepositoryImpl)GetListMembers(ctx context.Context)(entity.Members, error){
	query := SELECT_REAL_NAME.OrderBy("created_at ASC")
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR GetListMembers -> error: %v\n", err)
		return nil, err
	}
	prpd, err := mr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("ERROR GetListMembers -> error: %v\n", err)
		return nil, err
	}
	rows, err := prpd.Query()
	if err != nil {
		log.Printf("ERROR GetListMembers -> error: %v\n", err)
		return nil, err
	}

	members := entity.Members{}

	for rows.Next(){
		var member entity.Member

		err := rows.Scan(&member.Id, &member.RealName, &member.Birth, &member.Born, &member.Description, &member.CreatedAt, &member.UpdatedAt)
		if err != nil {
			log.Printf("ERROR GetListMembers -> error: %v\n", err)
			return nil, err
		}
		members = append(members, &member)
	}
	return members, nil
}