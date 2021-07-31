package domain_model_favorite

import (
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/domain/model/tweet"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/domain/model/user"
)

type Favorite struct {
	userId  string
	tweetId string
}

func newFavorite(userId string, tweetId string) *Favorite {
	return &Favorite{
		userId:  userId,
		tweetId: tweetId,
	}
}

func GenerateFavorite(user *User, tweet *Tweet) *Favorite {
	return newFavorite(user.Id(), tweet.Id())
}

func (fav Favorite) UserId() string {
	return fav.userId
}

func (fav Favorite) TweetId() string {
	return fav.tweetId
}
