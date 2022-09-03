package mongo_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"repo_mocktest/domain"
	userRepo "repo_mocktest/repository/mongo"
)

//reference
//https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/receiver/mongodbreceiver/client_test.go
//https://github.com/GoogleCloudPlatform/golang-samples/blob/main/compute/quickstart/compute_quickstart_sample_test.go

func TestUserRepository_Save(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		repo := userRepo.NewUserRepository(mt.DB, mt.Coll)
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		usr := domain.User{
			Name:    "John",
			Address: "Jakarta",
		}

		err := repo.Save(context.TODO(), &usr)
		assert.NoError(mt, err)
	})

	mt.Run("error", func(mt *mtest.T) {
		expectErr := "cannot transform type *domain.User to a BSON Document: WriteNull can only write while positioned on a Element or Value but is positioned on a TopLevel"
		repo := userRepo.NewUserRepository(mt.DB, mt.Coll)
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		var usr *domain.User //nil
		err := repo.Save(context.TODO(), usr)
		assert.Error(mt, err)
		assert.Equal(mt, expectErr, err.Error())
	})
}

func TestUserRepository_FindAll(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		repo := userRepo.NewUserRepository(mt.DB, mt.Coll)
		users := []domain.User{
			{
				Name:    "John",
				Address: "Jakarta",
			},
			{
				Name:    "Due",
				Address: "Bekasi",
			},
		}

		cursorResponses := make([]bson.D, 0)
		for i, user := range users {
			b, _ := json.Marshal(user)
			var d bson.D
			err := bson.UnmarshalExtJSON(b, true, &d)
			assert.NoError(mt, err)

			batch := mtest.NextBatch
			if i == 0 {
				batch = mtest.FirstBatch
			}
			curResponse := mtest.CreateCursorResponse(
				1,
				"DBName.CollectionName",
				batch,
				d,
			)
			cursorResponses = append(cursorResponses, curResponse)
		}

		killCursors := mtest.CreateCursorResponse(
			0,
			"DBName.CollectionName",
			mtest.NextBatch,
		)

		cursorResponses = append(cursorResponses, killCursors)
		mt.AddMockResponses(cursorResponses...)

		result, err := repo.FindAll(context.Background())

		assert.Nil(mt, err)
		assert.Equal(mt, 2, len(*result))
		assert.Equal(mt, users[0].Name, (*result)[0].Name)
		assert.Equal(mt, users[1].Name, (*result)[1].Name)
	})
}

//TODO make tests for all methods already implemented
