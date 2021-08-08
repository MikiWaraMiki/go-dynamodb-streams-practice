package infra_repository

import (
	"time"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/user"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Conn *gorm.DB
}

type UserDto struct {
	UUID      string    `gorm:"column:uuid"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func NewUserRepository(conn *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		Conn: conn,
	}
}

func (ur *UserRepositoryImpl) FindById(id *UserID) (*User, error) {
	var userDto UserDto

	if err := ur.Conn.Table("users").Where("uuid = ?", id.Value()).Find(&userDto).Error; err != nil {
		return nil, err
	}

	user := NewUser(id)
	return user, nil
}
