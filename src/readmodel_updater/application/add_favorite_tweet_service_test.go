package application

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/favorite_tweet"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/tweet"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

/**
	Generate Mock
**/
type MockUserRepository struct {
	mock.Mock
}

func (mock *MockUserRepository) FindById(id *UserID) (*User, error) {
	fmt.Println("Mock呼び出し")
	result := mock.Called(id)
	return result.Get(0).(*User), result.Error(1)
}

type MockTweetRepository struct {
	mock.Mock
}

func (mock *MockTweetRepository) FindById(tweetId *TweetID) (*Tweet, error) {
	fmt.Println("Mock呼び出し")
	result := mock.Called(tweetId)
	return result.Get(0).(*Tweet), result.Error(1)
}

type MockFavoriteTweetRepository struct {
	mock.Mock
}

func (mock *MockFavoriteTweetRepository) InsertOne(favoriteTweet *FavoriteTweet) error {
	fmt.Println("Mock呼び出し")
	result := mock.Called(favoriteTweet)

	return result.Error(0)
}

type AddFavoriteTweetServiceTestSuite struct {
	suite.Suite
	userRepository          UserRepository
	tweetRepository         TweetRepository
	favoriteTweetRepository FavoriteTweetRepository
}

func TestAddFavoriteTweetServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AddFavoriteTweetServiceTestSuite))
}

func (suite *AddFavoriteTweetServiceTestSuite) SetupTest() {
	// 正常系でセットアップしておく
	uuid, _ := uuid.NewRandom()

	userId, _ := NewUserID(uuid.String())
	returnedUser := NewUser(userId)

	tweetId, _ := NewTweetId(uuid.String())
	returnerTweet := NewTweet(tweetId, "hogehoge")

	// Mock
	mockUserRepo := new(MockUserRepository)
	mockUserRepo.On("FindById", mock.Anything).Return(returnedUser, nil)

	mockTweetRepo := new(MockTweetRepository)
	mockTweetRepo.On("FindById", mock.Anything).Return(returnerTweet, nil)

	mockFavoriteTweetRepo := new(MockFavoriteTweetRepository)
	mockFavoriteTweetRepo.On("InsertOne", mock.Anything).Return(nil)

	suite.userRepository = mockUserRepo
	suite.tweetRepository = mockTweetRepo
	suite.favoriteTweetRepository = mockFavoriteTweetRepo
}

func (suite *AddFavoriteTweetServiceTestSuite) TestAddFavoriteTweet() {
	suite.Run("指定したIDのユーザが存在しな場合はエラーを返す", func() {
		uuid, _ := uuid.NewRandom()

		mockUserRepo := new(MockUserRepository)
		mockUserRepo.On("FindById", mock.Anything).Return(&User{}, gorm.ErrRecordNotFound)

		service := NewAddFavoriteTweetService(
			mockUserRepo,
			suite.tweetRepository,
			suite.favoriteTweetRepository,
		)

		err := service.AddFavoriteTweet(uuid.String(), uuid.String())

		assert.NotNil(suite.T(), err)
		assert.Equal(suite.T(), "ユーザが存在しませんでした。", fmt.Sprintf("%v", err))
	})
	suite.Run("指定したIDのツイートが存在しない場合はエラーを返す", func() {
		uuid, _ := uuid.NewRandom()

		mockTweetRepo := new(MockTweetRepository)
		mockTweetRepo.On("FindById", mock.Anything).Return(&Tweet{}, gorm.ErrRecordNotFound)

		service := NewAddFavoriteTweetService(
			suite.userRepository,
			mockTweetRepo,
			suite.favoriteTweetRepository,
		)

		err := service.AddFavoriteTweet(uuid.String(), uuid.String())

		assert.NotNil(suite.T(), err)
		assert.Equal(suite.T(), "お気に入り対象のツイートが存在しませんでした。", fmt.Sprintf("%v", err))
	})
	suite.Run("お気に入り対象への新規追加処理が失敗した場合はエラーを返す", func() {
		uuid, _ := uuid.NewRandom()

		mockFavoriteTweetRepository := new(MockFavoriteTweetRepository)
		mockFavoriteTweetRepository.On("InsertOne", mock.Anything).Return(errors.New("invalid"))

		service := NewAddFavoriteTweetService(
			suite.userRepository,
			suite.tweetRepository,
			mockFavoriteTweetRepository,
		)

		err := service.AddFavoriteTweet(uuid.String(), uuid.String())

		assert.NotNil(suite.T(), err)
		assert.Equal(suite.T(), "お気に入り対象への追加に失敗しました。", fmt.Sprintf("%v", err))
	})
}
