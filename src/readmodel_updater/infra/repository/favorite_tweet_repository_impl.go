package infra_repository

import (
	"time"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/favorite_tweet"
	"gorm.io/gorm"
)

type FavoriteTweetRepositoryImpl struct {
	Conn *gorm.DB
}

type FavoriteTweetDto struct {
	ID        int        `gorm:"column:id"`
	UserUUID  string     `gorm:"column:user_uuid"`
	TweetID   string     `gorm:"column:tweet_id"`
	Content   string     `gorm:"column:content"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func NewFavoriteTweetRepository(conn *gorm.DB) *FavoriteTweetRepositoryImpl {
	return &FavoriteTweetRepositoryImpl{
		Conn: conn,
	}
}

func (repo *FavoriteTweetRepositoryImpl) InsertOne(favoriteTweet *FavoriteTweet) error {
	favoriteTweetDto := FavoriteTweetDto{
		UserUUID:  favoriteTweet.UserId(),
		TweetID:   favoriteTweet.TweetId(),
		Content:   favoriteTweet.Content(),
		CreatedAt: nil,
		UpdatedAt: nil,
	}

	err := repo.Conn.Table("favorite_tweets").
		Select("UserUUID", "TweetID", "Content").
		Create(&favoriteTweetDto).Error

	if err != nil {
		return err
	}

	return nil
}
