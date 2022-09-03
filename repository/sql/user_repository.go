package sql

import (
	"context"
	"database/sql"
	"errors"

	"repo_mocktest/domain"
)

type userRepository struct {
	db *sql.DB
}

//NewUserRepository implement UserRepository Interface
func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) Save(ctx context.Context, user *domain.User) error {
	q := "INSERT INTO users VALUES (?, ?, ?)"
	res, err := u.db.ExecContext(ctx, q, nil, user.Name, user.Address) //TODO mock with autoincrement
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows < 1 {
		return errors.New("no rows affected")
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = lastId
	return nil
}

func (u userRepository) UpdateByID(ctx context.Context, id string, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindAll(ctx context.Context, user *domain.User) (*[]*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) FindByID(ctx context.Context, user *domain.User) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) DeleteByID(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
