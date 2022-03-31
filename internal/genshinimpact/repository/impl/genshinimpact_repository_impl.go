package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/genshinimpact/entity"
	sq "github.com/Masterminds/squirrel"
)

type genshinimpactRepositoryImpl struct {
	DB *sql.DB
}

func ProvideGenshinImpactRepository(DB *sql.DB) *genshinimpactRepositoryImpl{
	return &genshinimpactRepositoryImpl{DB: DB}
}

var(
	SELECT_GENSHIN_IMPACT_MEMBERS = sq.Select("genshin_impact.id", "genshin_impact.member_id", "genshin_impact.ingame_name", "member.real_name", "genshin_impact.image", "genshin_impact.nationality", "genshin_impact.created_at", "genshin_impact.updated_at").From("genshin_impact").Join("member ON member.id = genshin_impact.member_id")
	SELECT_REAL_NAME = sq.Select("real_name").From("member")
	CHECK_GENSHIN_MEMBER = sq.Select("count(*)").From("genshin_impact")
)

func(gr genshinimpactRepositoryImpl)GetAllMembers(ctx context.Context)(entity.GenshinImpactMembers, entity.Members, error){
	query := SELECT_GENSHIN_IMPACT_MEMBERS.OrderBy("created_at ASC")
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}
	prpd, err := gr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}
	rows, err := prpd.Query()
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}
	genshinMembers := entity.GenshinImpactMembers{}
	members := entity.Members{}

	for rows.Next(){
		var genshinMember entity.GenshinImpactMember
		var member entity.Member

		err := rows.Scan(&genshinMember.Id, &genshinMember.MemberId, &genshinMember.InGameName, &member.RealName, &genshinMember.Image, &genshinMember.Nationality, &genshinMember.CreatedAt, &genshinMember.UpdatedAt)
		if err != nil {
			log.Printf("ERROR GetAllMembers -> error: %v\n", err)
			return nil, nil, err
		}
		genshinMembers = append(genshinMembers, &genshinMember)
		members = append(members, &member)
	}
	return genshinMembers, members, nil
}

func(gr genshinimpactRepositoryImpl)CheckMembersExists(ctx context.Context)(int, error){
	query := CHECK_GENSHIN_MEMBER
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR CheckMembersExists -> error: %v\n", err)
		return 0, err
	}
	prpd, err := gr.DB.PrepareContext(ctx, stmt)
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

func(gr genshinimpactRepositoryImpl)GetListMembers(ctx context.Context)(entity.Members, error){
	query := SELECT_REAL_NAME
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR GetListMembers -> error: %v\n", err)
		return nil, err
	}
	prpd, err := gr.DB.PrepareContext(ctx, stmt)
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