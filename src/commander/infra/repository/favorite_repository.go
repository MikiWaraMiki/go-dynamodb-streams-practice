package infra_repository

import (
	"errors"
	"fmt"
	"time"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/application/event"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/guregu/dynamo"
)

type FavoriteEventRepositoryImpl struct {
	db *dynamo.DB
}

func NewFavoriteEventRepository(db *dynamo.DB) *FavoriteEventRepositoryImpl {
	return &FavoriteEventRepositoryImpl{
		db: db,
	}
}

type ProviderTableItem struct {
	EventProviderId string `dynamo:"eventProviderId"`
	EventType       string `dynamo:"eventType"`
	Version         int    `dynamo:"version"`
}

type EventStoreItem struct {
	EventProviderId string            `dynamo:"eventProviderId"`
	Version         int               `dynamo:"version"`
	CreatedAt       time.Time         `dynamo:"createdAt"`
	Body            map[string]string `dynamo:"body"`
}

func (repo FavoriteEventRepositoryImpl) Store(event AddTweetFavoriteEvent) error {
	var err error

	if err = repo.CreateProvider(event.UserId); err != nil {
		return err
	}

	providerItem := repo.GetProviderData(event.UserId)
	if providerItem == nil {
		return errors.New("not found provider data")
	}

	providerTable := repo.db.Table("provider-store")
	eventTable := repo.db.Table("event-store")

	body := map[string]string{}
	body["eventId"] = "add-favorite"
	body["tweetId"] = event.TweetId
	eventStoreItem := EventStoreItem{
		EventProviderId: providerItem.EventProviderId,
		Version:         providerItem.Version + 1,
		CreatedAt:       time.Now(),
		Body:            body,
	}

	putEventOperation := eventTable.Put(&eventStoreItem)
	updateProviderOperation := providerTable.Update(
		"eventProviderId",
		event.UserId,
	).SetExpr(
		"'version' = 'version' + ?",
		1,
	).If(
		"version = ?",
		providerItem.Version,
	)

	err = repo.db.WriteTx().Put(putEventOperation).Update(updateProviderOperation).Run()
	if err != nil {
		fmt.Printf("failed transcation [%v]\n", err)
		return err
	}

	return nil
}

func (repo FavoriteEventRepositoryImpl) GetProviderData(userId string) *ProviderTableItem {
	providerTable := repo.db.Table("provider-store")

	var item ProviderTableItem

	providerTable.Get("eventProviderId", userId).Consistent(true).One(&item)

	return &item
}

func (repo FavoriteEventRepositoryImpl) CreateProvider(userId string) error {
	// Create Provider Table Record if not exists
	providerTable := repo.db.Table("provider-store")

	item := ProviderTableItem{
		EventProviderId: userId,
		EventType:       "add_favorite",
		Version:         0,
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
