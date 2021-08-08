package infra_repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/user"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepositoryImplTestSuite struct {
	suite.Suite
	userRepository *UserRepositoryImpl
	mock           sqlmock.Sqlmock
}

func (suite *UserRepositoryImplTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	mockConn, _ := gorm.Open(
		mysql.New(mysql.Config{
			Conn:                      db,
			SkipInitializeWithVersion: true,
		}),
		&gorm.Config{},
	)
	userRepository := NewUserRepository(mockConn)

	suite.userRepository = userRepository
}

func (suite *UserRepositoryImplTestSuite) TearDownTest() {
	db, _ := suite.userRepository.Conn.DB()
	db.Close()
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryImplTestSuite))
}

func (suite *UserRepositoryImplTestSuite) TestFindById() {
	suite.Run("ユーザーが存在する場合はエラーが発生しない", func() {
		uuid, _ := uuid.NewRandom()
		userId, _ := NewUserID(uuid.String())

		suite.mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM ` + "`users`" + ` WHERE uuid = ?`)).
			WithArgs(userId.Value()).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
				AddRow(userId.Value(), "hogehoge"))

		_, err := suite.userRepository.FindById(userId)

		require.NoError(suite.T(), err)
	})
}
