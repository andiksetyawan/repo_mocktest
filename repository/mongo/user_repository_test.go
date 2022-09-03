package mongo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"repo_mocktest/domain"
	userRepo "repo_mocktest/repository/mongo"
)

//reference
//https://github.com/zeromicro/go-zero/tree/master/core/stores/mon
//https://github.com/GoogleCloudPlatform/golang-samples/blob/main/compute/quickstart/compute_quickstart_sample_test.go
//https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/main/receiver/mongodbreceiver/client_test.go

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

//TODO make tests for all methods already implemented
