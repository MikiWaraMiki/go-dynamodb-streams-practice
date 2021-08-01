package application_event

import (
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/domain/model/tweet"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/domain/model/user"
)

type TweetFavoriteEventCommandProcessor struct {
}

func (command TweetFavoriteEventCommandProcessor) Execute(
	userId UserID,
	tweet Tweet,
) error {
	return nil
}
