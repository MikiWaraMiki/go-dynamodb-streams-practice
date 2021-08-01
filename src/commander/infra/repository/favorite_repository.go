package infra_repository

import (
	"fmt"
	"time"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/application/event"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/guregu/dynamo"
)

type FavoriteEventRepository struct {
	db *dynamo.DB
}

type ProviderTableItem struct {
	EventProviderId string    `dynamo:"eventProviderId"`
	EventType       string    `dynamo:"eventType"`
	Version         int       `dynamo:"version"`
	CreatedAt       time.Time `dynamo:"createdAt"`
}

func (repo FavoriteEventRepository) Store(event AddTweetFavoriteEvent) {
}

func (repo FavoriteEventRepository) GetProviderData(userId string) *ProviderTableItem {
	providerTable := repo.db.Table("provider-store")

	var item ProviderTableItem

	providerTable.Get("eventProviderId", userId).Consistent(true).One(&item)

	return &item
}

func (repo FavoriteEventRepository) CreateProvider(userId string) error {
	// Create Provider Table Record if not exists
	providerTable := repo.db.Table("provider-store")

	item := ProviderTableItem{
		EventProviderId: userId,
		EventType:       "add_favorite",
		Version:         0,
		CreatedAt:       time.Now(),
	}

	err := providerTable.Put(&item).If("attribute_not_exists(eventProviderId)").Run()

	if err != nil {
		if ae, ok := err.(awserr.RequestFailure); ok && ae.Code() == "ConditionalCheckFailedException" {
			fmt.Printf("already provider is exist\n")
			return nil
		}
		fmt.Printf("failed to put item[%v]\n", err)
		return err
	}

	return nil
}
