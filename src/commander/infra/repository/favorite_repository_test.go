package infra_repository

import (
	"testing"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/infra/handler"
	"github.com/guregu/dynamo"
)

func createDbSession() *dynamo.DB {
	return NewDynamoDBSession("local")
}

func createRepository() *FavoriteEventRepository {
	db := createDbSession()

	return &FavoriteEventRepository{
		db: db,
	}
}

func TestGetProviderData(t *testing.T) {
	t.Run("データを取得できること", func(t *testing.T) {
		repo := createRepository()

		repo.CreateProvider("user2")

		result := repo.GetProviderData("user2")

		if result == nil {
			t.Fatalf("expected data, result nil")
		}

		db := NewDynamoDBSession("local")
		db.Table("provider-store").Delete("eventProviderId", "user2").Run()
	})
}

func TestCreateProvider(t *testing.T) {
	t.Run("エラーが発生しないこと", func(t *testing.T) {
		repo := createRepository()

		err := repo.CreateProvider("user2")

		if err != nil {
			t.Fatalf("error happend %v", err)
		}

		db := NewDynamoDBSession("local")
		db.Table("provider-store").Delete("eventProviderId", "user2").Run()
	})
	t.Run("テーブルにデータが作成されていない場合は、データが作成されること", func(t *testing.T) {
		repo := createRepository()

		repo.CreateProvider("user2")

		db := NewDynamoDBSession("local")

		var item ProviderTableItem
		db.Table("provider-store").Get("eventProviderId", "user2").One(&item)

		if &item == nil {
			t.Fatalf("expected: getItem, result is nil")
		}

		db.Table("provider-store").Delete("eventProviderId", "user2").Run()
	})
}
