package sql_test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"repo_mocktest/domain"
	userRepo "repo_mocktest/repository/sql"
)

func NewDBTest() (db *sql.DB, mock sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return
}

func TestUserRepository_Save(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, mock := NewDBTest()
		defer db.Close()

		usr := &domain.User{
			Name:    "John",
			Address: "Tuban",
		}

		var lastInsertID int64 = 1
		expectResult := sqlmock.NewResult(lastInsertID, 1)
		mock.ExpectExec("INSERT INTO users VALUES (?, ?, ?)").WithArgs(nil, usr.Name, usr.Address).WillReturnResult(expectResult)

		repo := userRepo.NewUserRepository(db)
		err := repo.Save(context.TODO(), usr)
		assert.NoError(t, err)
		assert.Equal(t, lastInsertID, usr.ID)
	})

	// TODO error cases
}

//TODO make tests for all methods already implemented
