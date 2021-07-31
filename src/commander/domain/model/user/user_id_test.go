package domain_model_user

import "testing"

func EmptyUserIdTest(t *testing.T) {
	t.Run("空文字の場合はエラーが返されること", func(t *testing.T) {
		_, err := NewUserID("")

		if err != nil {
			t.Fatal("expected=(id is required error), result = nil")
		}
	})
}
