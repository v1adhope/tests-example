// TODO: squirrel must be removed?
package repository

import (
	"context"
	"fmt"

	"github.com/v1adhope/tests-example/internal/entity"
	"github.com/v1adhope/tests-example/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

type MemberRepo struct {
	*postgres.Postgres
}

func NewMemberRepo(pg *postgres.Postgres) *MemberRepo {
	return &MemberRepo{pg}
}

func (pg *MemberRepo) Insert(ctx context.Context, member *entity.Member) (*uint64, error) {
	sql, args, err := pg.Builder.Insert("members").
		Columns("first_name", "last_name", "age").
		Values(member.FirstName, member.LastName, member.Age).
		Suffix("RETURNING \"member_id\"").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("Repo: Insert: Insert: %s", err)
	}

	id := new(uint64)

	err = pg.Pool.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("Repo: Insert: QueryRow: %s", err)
	}

	return id, nil
}

func (pg *MemberRepo) GetByID(ctx context.Context, id uint64) (*entity.Member, error) {
	sql, args, err := pg.Builder.Select("first_name, last_name, age").
		From("members").
		Where(squirrel.Eq{"member_id": id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("Repo: GetByID: Select: %s", err)
	}

	memberDTO := &MemberDTO{}
	err = pg.Pool.QueryRow(ctx, sql, args...).Scan(&memberDTO.FirstName, &memberDTO.LastName, &memberDTO.Age)
	if err != nil {
		return nil, fmt.Errorf("Repo: GetByID: QueryRow: %s", err)
	}

	return memberDTO.ToEntity(), nil
}

// TODO: ID must be ID > 0
func (pg *MemberRepo) Update(ctx context.Context, member *entity.Member) error {
	buildS := pg.Builder.Update("members").Where(squirrel.Eq{"member_id": member.ID})

	if &member.FirstName != nil {
		buildS = buildS.Set("first_name", member.FirstName)
	}

	if &member.LastName != nil {
		buildS = buildS.Set("last_name", member.LastName)
	}

	if &member.Age != nil {
		buildS = buildS.Set("age", member.Age)
	}

	sql, args, err := buildS.ToSql()
	if err != nil {
		return fmt.Errorf("Repo: Update: ToSql: %s", err)
	}

	_, err = pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("Repo: Update: Exec: %s", err)
	}

	return nil
}

func (pg *MemberRepo) Delete(ctx context.Context, id *uint64) error {
	sql, args, err := pg.Builder.Delete("members").
		Where(squirrel.Eq{"member_id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("Repo: Delete: Delete: %s", err)
	}

	_, err = pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("Repo: Delete: Exec: %s", err)
	}

	return nil
}
