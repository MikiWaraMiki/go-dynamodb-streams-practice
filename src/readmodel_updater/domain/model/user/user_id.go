package domain_model_user

import (
	"errors"

	"github.com/google/uuid"
)

type UserID struct {
	value string
}

func NewUserID(id string) (*UserID, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	if !isValidateUUID(id) {
		return nil, errors.New("id is invalid")
	}
	// TODO: UUIDのバリデーション
	return &UserID{
		value: id,
	}, nil
}

func isValidateUUID(id string) bool {
	_, err := uuid.Parse(id)

	return err == nil
}

func (id *UserID) Value() string {
	return id.value
}
