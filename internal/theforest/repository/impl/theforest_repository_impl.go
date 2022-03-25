package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/theforest/entity"
	sq "github.com/Masterminds/squirrel"
)

type theforestRepositoryImpl struct {
	DB *sql.DB
}

func ProvideTheForestRepository(DB *sql.DB) *theforestRepositoryImpl{
	return &theforestRepositoryImpl{DB: DB}
}

var(
	SELECT_THE_FOREST_MEMBERS = sq.Select("the_forest.id", "the_forest.member_id", "the_forest.ingame_name", "member.real_name", "the_forest.image", "the_forest.nationality", "the_forest.created_at", "the_forest.updated_at").From("the_forest").Join("member ON member.id = the_forest.member_id")
	SELECT_REAL_NAME =sq.Select("real_name").From("member")
	CHECK_THE_FOREST_MEMBERS = sq.Select("count(*)").From("the_forest")
)

func(fr theforestRepositoryImpl)GetAllMembers(ctx context.Context)(entity.TheForestMembers, entity.Members, error){
	query := SELECT_THE_FOREST_MEMBERS.OrderBy("created_at ASC")
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}
	prpd, err := fr.DB.PrepareContext(ctx, stmt)
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}
	rows, err := prpd.Query()
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, nil, err
	}
	membersTheForest := entity.TheForestMembers{}
	members := entity.Members{}
	for rows.Next(){
		var memberTheForest entity.TheForestMember
		var member entity.Member

		err := rows.Scan(&memberTheForest.Id, &memberTheForest.MemberId, &memberTheForest.InGameName, &member.RealName, &memberTheForest.Image, &memberTheForest.Nationality, &memberTheForest.CreatedAt, &memberTheForest.UpdatedAt)
		if err != nil {
			log.Printf("ERROR GetAllMembers -> error: %v\n", err)
			return nil, nil, err
		}
		membersTheForest = append(membersTheForest, &memberTheForest)
		members = append(members, &member)
	}
	return membersTheForest, members, nil
}

func(fr theforestRepositoryImpl)CheckMembersExists(ctx context.Context)(int, error){
	query := CHECK_THE_FOREST_MEMBERS
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR CheckMembersExists -> error: %v\n", err)
		return 0, err
	}
	prpd, err := fr.DB.PrepareContext(ctx, stmt)
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

func(fr theforestRepositoryImpl)GetListMembers(ctx context.Context)(entity.Members, error){
	query := SELECT_REAL_NAME.OrderBy("created_at ASC")
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR GetListMembers -> error: %v\n", err)
		return nil, err
	}
	prpd, err := fr.DB.PrepareContext(ctx, stmt)
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