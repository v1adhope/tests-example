package repository

import (
	"context"
	"polygon/internal/entity"
	"polygon/pkg/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

const _PG_URL = "postgres://postgres:secret@127.0.0.1:5432/test"
const _COMPOSE_PATH = "../../../compose.yaml"

func TestRepo(t *testing.T) {
	compose, err := tc.NewDockerCompose(_COMPOSE_PATH)
	assert.NoError(t, err, "NewDockerComposeAPI()")
	t.Cleanup(func() {
		assert.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
	})
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	assert.NoError(t, compose.Up(ctx, tc.Wait(true)), "compose.Up()")
	member := entity.Member{
		FirstName: "Vlad",
		LastName:  "Gardner",
		Age:       21,
	}
	pg, err := postgres.Build(context.Background(), _PG_URL)
	if err != nil {
		t.Error(err)
	}
	repo := NewMemberRepo(pg)

	SUT, err := repo.Insert(context.Background(), &member)

	assert.Equal(t, uint64(1), *SUT)
	assert.Equal(t, nil, err)

	//TODO: decomp

	SUT2, err := repo.GetByID(context.Background(), 1)

	assert.Equal(t, member, *SUT2)
	assert.Equal(t, nil, err)

	//TODO: decomp

	updateMember := entity.Member{
		ID:        1,
		FirstName: "Not Vlad",
		LastName:  "Not Gardner",
		Age:       42,
	}

	SUT3 := repo.Update(context.Background(), &updateMember)

	assert.NoError(t, SUT3)

	//TODO: decomp

	memberID := uint64(1)

	SUT4 := repo.Delete(context.Background(), &memberID)

	assert.NoError(t, SUT4)
}
