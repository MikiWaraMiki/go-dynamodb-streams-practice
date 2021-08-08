package main

import (
	"os"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/application/event"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/application/usecase"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/infra/handler"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/infra/repository"
	"github.com/aws/aws-lambda-go/lambda"
)

type FavoriteEvent struct {
	TweetID string `json:tweetID`
	UserID  string `json:userID`
}

type FavoriteResponse struct {
	Message string `json:message`
}

func handler(event FavoriteEvent) (*FavoriteResponse, error) {
	environment := os.Getenv("ENVIRONMENT")
	dbSession := NewDynamoDBSession(environment)
	repository := NewFavoriteEventRepository(dbSession)

	commander := NewTweetFavoriteEventCommandProcessor(repository)

	service := NewTweetFavoriteService(commander)

	err := service.AddFavorite(event.TweetID, event.UserID)

	if err != nil {
		return nil, err
	}

	return &FavoriteResponse{
		Message: "success",
	}, nil
}

func main() {
	lambda.Start(handler)
}
