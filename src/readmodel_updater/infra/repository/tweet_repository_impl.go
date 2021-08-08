package infra_repository

import (
	"time"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/tweet"
	"gorm.io/gorm"
)

type TweetRepositoryImpl struct {
	Conn *gorm.DB
}

type TweetDTO struct {
	ID        string    `gorm:"column:id`
	UserUUID  string    `gorm:"column:user_uuid"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func NewTweetRepository(conn *gorm.DB) *TweetRepositoryImpl {
	return &TweetRepositoryImpl{
		Conn: conn,
	}
}

func (tr *TweetRepositoryImpl) FindById(id *TweetID) (*Tweet, error) {
	var tweetDto TweetDTO

	if err := tr.Conn.Table("tweets").Where("id = ?", id.Value()).Find(&tweetDto).Error; err != nil {
		return nil, err
	}

	return NewTweet(id, tweetDto.Content), nil
}
