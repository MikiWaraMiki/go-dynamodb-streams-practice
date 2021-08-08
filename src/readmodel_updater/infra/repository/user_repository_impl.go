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
	gorm.Model

	Id        string
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewUserRepository(conn *gorm.db) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		Conn: conn,
	}
}

func (ur *UserRepositoryImpl) findById(id *UserID) (*User, error) {
	userDto := UserDto{
		Id: id.Value(),
	}

	if err := ur.Conn.First(&userDto).Error; err != nil {
		return nil, error
	}

	return NewUser(
		NewUserId(userDto.Id),
	), nil
}
