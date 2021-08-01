package domain_model_user

import (
	"errors"
	"fmt"

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

func GenerateUserId() (*UserID, error) {
	id, err := uuid.NewRandom()

	if err != nil {
		fmt.Printf("failed generate uuid")
		return nil, err
	}

	return NewUserID(id.String())
}

func isValidateUUID(id string) bool {
	_, err := uuid.Parse(id)

	return err == nil
}
