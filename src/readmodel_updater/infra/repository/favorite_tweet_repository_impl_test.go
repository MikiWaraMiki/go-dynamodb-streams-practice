package infra_repository

import (
	"database/sql/driver"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/favorite_tweet"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/tweet"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/user"
)

// NOTE: https://github.com/DATA-DOG/go-sqlmock/blob/b9ca56ce96879f5362120ae10866bbf66f2c5db6/argument_test.go
type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type FavoriteTweetRepositoryImplSuite struct {
	suite.Suite
	favoriteTweetRepository *FavoriteTweetRepositoryImpl
	mock                    sqlmock.Sqlmock
}

func (suite *FavoriteTweetRepositoryImplSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	mockConn, _ := gorm.Open(
		mysql.New(mysql.Config{
			Conn:                      db,
			SkipInitializeWithVersion: true,
		}),
		&gorm.Config{},
	)

	favoriteTweetRepository := NewFavoriteTweetRepository(mockConn)

	suite.favoriteTweetRepository = favoriteTweetRepository
}

func (suite *FavoriteTweetRepositoryImplSuite) TearDownTest() {
	db, _ := suite.favoriteTweetRepository.Conn.DB()
	db.Close()
}

func TestFavoriteTweetRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(FavoriteTweetRepositoryImplSuite))
}

func (suite *FavoriteTweetRepositoryImplSuite) TestInsertOnce() {
	suite.Run("異常値が無い場合は、エラーが返されないこと", func() {
		userIdRaw, _ := uuid.NewRandom()
		userId, _ := NewUserID(userIdRaw.String())
		user := NewUser(userId)

		tweetIdRaw, _ := uuid.NewRandom()
		tweetId, _ := NewTweetId(tweetIdRaw.String())
		content := strings.Repeat("a", 100)
		tweet := NewTweet(tweetId, content)

		favoriteTweet := NewFavoriteTweet(user, tweet)

		suite.mock.ExpectBegin()
		suite.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `favorite_tweets`")).
			WithArgs(userIdRaw, tweetIdRaw, content, AnyTime{}, AnyTime{}).
			WillReturnResult(sqlmock.NewResult(1, 1))
		suite.mock.ExpectCommit()

		err := suite.favoriteTweetRepository.InsertOne(favoriteTweet)

		assert.Nil(suite.T(), err)
	})
	suite.Run("INSERT時にエラーが発生した場合は、エラーを返す", func() {
		userIdRaw, _ := uuid.NewRandom()
		userId, _ := NewUserID(userIdRaw.String())
		user := NewUser(userId)

		tweetIdRaw, _ := uuid.NewRandom()
		tweetId, _ := NewTweetId(tweetIdRaw.String())
		content := strings.Repeat("a", 100)
		tweet := NewTweet(tweetId, content)

		favoriteTweet := NewFavoriteTweet(user, tweet)

		suite.mock.ExpectBegin()
		suite.mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `favorite_tweets`")).
			WithArgs(userIdRaw, tweetIdRaw, content, AnyTime{}, AnyTime{}).
			WillReturnError(gorm.ErrRegistered)
		suite.mock.ExpectRollback()

		err := suite.favoriteTweetRepository.InsertOne(favoriteTweet)

		assert.NotNil(suite.T(), err)
	})
}
