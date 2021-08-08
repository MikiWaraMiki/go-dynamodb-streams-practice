package infra_repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/tweet"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TweetRepositoryImplTestSuite struct {
	suite.Suite
	tweetRepository *TweetRepositoryImpl
	mock            sqlmock.Sqlmock
}

func (suite *TweetRepositoryImplTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	mockConn, _ := gorm.Open(
		mysql.New(mysql.Config{
			Conn:                      db,
			SkipInitializeWithVersion: true,
		}),
		&gorm.Config{},
	)
	tweetRepository := NewTweetRepository(mockConn)

	suite.tweetRepository = tweetRepository
}

func (suite *TweetRepositoryImplTestSuite) TearDownTest() {
	db, _ := suite.tweetRepository.Conn.DB()
	db.Close()
}

func TestTweetRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TweetRepositoryImplTestSuite))
}

func (suite *TweetRepositoryImplTestSuite) TestFindById() {
	suite.Run("ツイートが存在する場合はツイートオブジェクトを返すこと", func() {
		id, _ := uuid.NewRandom()
		userId, _ := uuid.NewRandom()
		tweetId, _ := NewTweetId(id.String())
		content := "hgoehogehoge"

		suite.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM ` + "`tweets` " + `WHERE id = ?`)).
			WithArgs(tweetId.Value()).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_uuid", "content"}).
				AddRow(id, userId.String(), content))

		tweet, err := suite.tweetRepository.FindById(tweetId)

		assert.Nil(suite.T(), err)

		assert.Equal(suite.T(), tweetId.Value(), tweet.Id())
		assert.Equal(suite.T(), content, tweet.Content())
	})
	suite.Run("ツイートが存在しない場合はエラーを返すこと", func() {
		id, _ := uuid.NewRandom()
		tweetId, _ := NewTweetId(id.String())

		suite.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM ` + "`tweets` " + `WHERE id = ?`)).
			WithArgs(tweetId.Value()).
			WillReturnError(gorm.ErrRecordNotFound)

		_, err := suite.tweetRepository.FindById(tweetId)

		assert.NotNil(suite.T(), err)
	})
}
