package domain_model_favorite_tweet

import (
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/tweet"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/user"
)

type FavoriteTweet struct {
	user  *User
	tweet *Tweet
}

func NewFavoriteTweet(user *User, tweet *Tweet) *FavoriteTweet {
	return &FavoriteTweet{
		user:  user,
		tweet: tweet,
	}
}

func (ft *FavoriteTweet) UserId() string {
	return user.Id()
}

func (ft *FavoriteTweet) TweetId() string {
	return tweet.Id()
}

func (ft *FavoriteTweet) Content() string {
	return tweet.Content()
}
