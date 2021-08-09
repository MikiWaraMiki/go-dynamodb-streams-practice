package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/lambda"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/application"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/handler"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/infra/repository"
)

func handler(event events.DynamoDBEvent) {
	connection := NewGormConnection()

	userRepository := NewUserRepository(connection)
	tweetRepository := NewTweetRepository(connection)
	favoriteTweetRepository := NewFavoriteTweetRepository(conn)

	addFavoriteTweetService := NewAddFavoriteTweetService(
		userRepository,
		tweetRepository,
		favoriteTweetRepository,
	)

	for _, record := range event.Records {
		fmt.Printf("Processing request data for event ID %s, type %s.\n", record.EventID, record.EventName)

		eventId := record.Change.NewImage["body"]["eventId"].String()
		tweetId := record.Change.NewImage["body"]["tweetId"].String()
		userId := record.Change.NewImage["body"]["userId"].String()

		fmt.Printf("eventId: %s, userId: %s, userId: %s \n", eventId, tweetId, userId)

		if eventId == "add-favorite" {
			addFavoriteTweetService.AddFavoriteTweet(
				userId,
				tweetId,
			)
		}
	}
}

func main() {
	lambda.Start(handler)
}
