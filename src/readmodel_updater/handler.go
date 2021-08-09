package main

import (
	"fmt"
	"os"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/application"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/handler"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/infra/repository"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rs/zerolog"
)

func handler(event events.DynamoDBEvent) {
	fmt.Printf("Event Received %v", event)
	connection := NewGormConnection()

	userRepository := NewUserRepository(connection)
	tweetRepository := NewTweetRepository(connection)
	favoriteTweetRepository := NewFavoriteTweetRepository(connection)

	addFavoriteTweetService := NewAddFavoriteTweetService(
		userRepository,
		tweetRepository,
		favoriteTweetRepository,
	)
	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().
		Timestamp().
		Str("role", "logger-lambda").
		Logger()
	logger.Info().Interface("event", event).Send()

	for _, record := range event.Records {
		fmt.Printf("Processing request data for event ID %s, type %s.\n", record.EventID, record.EventName)

		body := record.Change.NewImage["body"].Map()
		eventId := body["eventId"].String()
		tweetId := body["tweetId"].String()
		userId := body["userId"].String()

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
