package application_usecase

import (
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/application/event"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/domain/model/tweet"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/domain/model/user"
)

type TweetFavoriteService struct {
	command TweetFavoriteEventCommandProcessor
}

func (sv TweetFavoriteService) AddFavorite(tweetIdStr string, userIdStr string) error {
	tweetId, err := NewTweetId(tweetIdStr)
	if err != nil {
		return err
	}
	tweet := NewTweet(tweetId)

	// NOTE: デモなのでUserIDは都度生成
	userId, err := GenerateUserId()
	if err != nil {
		return err
	}
	user := NewUser(userId)

	if err := sv.command.Execute(user, tweet); err != nil {
		return err
	}

	return nil
}
