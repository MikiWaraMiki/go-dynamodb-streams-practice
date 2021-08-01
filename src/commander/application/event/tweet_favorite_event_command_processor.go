package application_event

import (
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/domain/model/tweet"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/domain/model/user"
)

type TweetFavoriteEventCommandProcessor struct {
	favoriteRepository FavoriteRepository
}

func (command TweetFavoriteEventCommandProcessor) Execute(
	user *User,
	tweet *Tweet,
) error {
	event := AddTweetFavoriteEvent{
		UserId:  user.Id(),
		TweetId: tweet.Id(),
	}

	if err := command.favoriteRepository.Store(event); err != nil {
		return err
	}

	return nil
}
