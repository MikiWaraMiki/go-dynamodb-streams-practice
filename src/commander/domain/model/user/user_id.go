package domain_model_user

import "errors"

type UserID struct {
	value string
}

func NewUserID(id string) (*UserID, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	// TODO: UUIDのバリデーション
	return &UserID{
		value: id,
	}, nil
}
