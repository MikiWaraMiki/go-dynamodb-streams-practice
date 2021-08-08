package favorite_tweet

import (
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/tweet"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/user"
)

type FavoriteTweet struct {
	user_id  string
	tweet_id int
	body     string
}

func NewFavoriteTweet(user_id *UserID, tweet *Tweet) *FavoriteTweet {
	return &FavoriteTweet{
		user_id:  user_id.Value(),
		tweet_id: tweet.Id(),
		body:     tweet.Body(),
	}
}
