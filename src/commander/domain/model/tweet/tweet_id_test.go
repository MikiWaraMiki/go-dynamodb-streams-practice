package domain_model_tweet

import "testing"

func TestEmptyTweetId(t *testing.T) {
	t.Run("空文字の場合はエラーが発生すること", func(t *testing.T) {
		_, err := NewTweetId("")

		if err == nil {
			t.Fatal("expected=(id is required), result=nil")
		}
	})
}
