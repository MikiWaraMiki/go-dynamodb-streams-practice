package infra_repository

import (
	"testing"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/infra/handler"
)

func createRepository() *FavoriteEventRepository {
	db := NewDynamoDBSession("local")

	return &FavoriteEventRepository{
		db: db,
	}
}

func TestCreateProvider(t *testing.T) {
	t.Run("エラーが発生しないこと", func(t *testing.T) {
		repo := createRepository()

		err := repo.CreateProvider("user2")

		if err != nil {
			t.Fatalf("error happend %v", err)
		}
	})

	db := NewDynamoDBSession("local")

	db.Table("provider-store").Delete("eventProviderId", "user2")
}
