package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/Jack-Gang-Worldwide/backend-web/internal/members/entity"
	sq "github.com/Masterminds/squirrel"
)

type membersRepositoryImpl struct {
	DB *sql.DB
}

var (
	SELECT_MEMBERS = sq.Select("id", "real_name", "birth", "born", "description", "created_at", "updated_at").From("member")
	CHECK_MEMBERS = sq.Select("count(*)").From("member")
)

func ProvideMemberRepository(DB *sql.DB)*membersRepositoryImpl{
	return &membersRepositoryImpl{DB: DB}
}

func(m membersRepositoryImpl)GetAllMembers(ctx context.Context)(entity.Members, error){
	query := SELECT_MEMBERS.OrderBy("created_at ASC")
	stmt, _, err := query.ToSql()
	if err != nil {
		log.Printf("ERROR GetAllMembers -> error: %v\n", err)
		return nil, err
	}
	prpd, err := m.DB.PrepareContext(ctx, stmt)
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

func(m membersRepositoryImpl)CheckMembersExists(ctx context.Context)(int, error){
	stmt, _, err := CHECK_MEMBERS.ToSql()
	if err != nil {
		log.Printf("ERROR CheckMembersExists -> error: %v\n", err)
		return 0, err
	}
	prpd, err := m.DB.PrepareContext(ctx, stmt)
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