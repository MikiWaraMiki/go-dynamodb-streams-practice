package domain_model_user

import "testing"

func TestEmptyUserId(t *testing.T) {
	t.Run("空文字の場合はエラーが返されること", func(t *testing.T) {
		_, err := NewUserID("")

		if err == nil {
			t.Fatal("expected=(id is required error), result = nil")
		}
	})
}

func TestInvalidUserId(t *testing.T) {
	t.Run("UUID以外が指定されている場合はエラーが返されること", func(t *testing.T) {
		_, err := NewUserID("aaaaaaaa")

		if err == nil {
			t.Fatal("expected=(id is invalid), result = nil")
		}
	})
}
