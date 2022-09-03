package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"repo_mocktest/domain"
)

type userRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

//NewUserRepository implement UserRepository Interface
func NewUserRepository(db *mongo.Database, coll *mongo.Collection) domain.UserRepository {
	return &userRepository{db: db, coll: coll}
}

func (u *userRepository) Save(ctx context.Context, user *domain.User) error {
	res, err := u.coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	log.Printf("last inserted objectID: %s", res.InsertedID)
	return nil
}

func (u *userRepository) UpdateByID(ctx context.Context, id string, user *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindAll(ctx context.Context) (*[]domain.User, error) {
	cur, err := u.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	users := make([]domain.User, 0)
	for cur.Next(ctx) {
		var usr domain.User
		err := cur.Decode(&usr)
		if err != nil {
			log.Println(err)
			continue
		}

		users = append(users, usr)
	}

	return &users, nil
}

func (u *userRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) DeleteByID(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
