package domain

import "context"

//User Domain
type User struct {
	ID      int64  `json:"id" bson:"id"`
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
}

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	UpdateByID(ctx context.Context, id string, user *User) error
	FindAll(ctx context.Context) (*[]User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	DeleteByID(ctx context.Context, id string) error
}
