package application

import (
	"errors"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/favorite_tweet"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/tweet"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/user"
)

type AddFavoriteTweetService struct {
	userRepository          UserRepository
	tweetRepository         TweetRepository
	favoriteTweetRepository FavoriteTweetRepository
}

func NewAddFavoriteTweetService(userRepository UserRepository, tweetRepository TweetRepository, favoriteTweetRepository FavoriteTweetRepository) *AddFavoriteTweetService {
	return &AddFavoriteTweetService{
		userRepository:          userRepository,
		tweetRepository:         tweetRepository,
		favoriteTweetRepository: favoriteTweetRepository,
	}
}

func (service *AddFavoriteTweetService) AddFavoriteTweet(addingUserId string, targetTweetId string) error {
	userId, err := NewUserID(addingUserId)
	if err != nil {
		return err
	}

	user, err := service.userRepository.FindById(userId)
	if err != nil {
		return errors.New("ユーザが存在しませんでした。")
	}

	tweetId, err := NewTweetId(targetTweetId)
	if err != nil {
		return err
	}
	tweet, err := service.tweetRepository.FindById(tweetId)
	if err != nil {
		return errors.New("お気に入り対象のツイートが存在しませんでした。")
	}

	favoriteTweet := NewFavoriteTweet(user, tweet)
	err = service.favoriteTweetRepository.InsertOne(favoriteTweet)

	if err != nil {
		return errors.New("お気に入り対象への追加に失敗しました。")
	}

	return nil
}
