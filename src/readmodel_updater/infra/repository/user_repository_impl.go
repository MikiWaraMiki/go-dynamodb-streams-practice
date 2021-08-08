package infra_repository

import (
	"errors"
	"time"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/readmodel_updater/domain/model/user"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Conn *gorm.DB
}

type UserDto struct {
	gorm.Model

	Id        string
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewUserRepository(conn *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		Conn: conn,
	}
}

func (ur *UserRepositoryImpl) FindById(id *UserID) (*User, error) {
	userDto := UserDto{
		Id: id.Value(),
	}

	if err := ur.Conn.First(&userDto).Error; err != nil {
		return nil, errors.New("user is not found")
	}

	userId, _ := NewUserID(userDto.Id)
	return NewUser(userId)
}
