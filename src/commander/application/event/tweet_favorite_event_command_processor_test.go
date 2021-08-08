package application_event

import (
	"errors"
	"testing"

	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/domain/model/tweet"
	. "github.com/MikiWaraMiki/go-dynamodb-streams-practice/src/commander/domain/model/user"
)

type MockSuccessCaseFavoriteRepository struct{}

func (repo MockSuccessCaseFavoriteRepository) Store(event AddTweetFavoriteEvent) error {
	return nil
}

type MockFailedCaseFavoriteRepository struct{}

func (repo MockFailedCaseFavoriteRepository) Store(event AddTweetFavoriteEvent) error {
	return errors.New("failed store")
}

func SetUp() (*User, *Tweet, error) {
	userId, error := GenerateUserId()
	if error != nil {
		return nil, nil, error
	}
	user := NewUser(userId)

	tweetId, error := NewTweetId("hogehoge")
	if error != nil {
		return nil, nil, error
	}
	tweet := NewTweet(tweetId)

	return user, tweet, nil
}

func TestExecute(t *testing.T) {
	t.Run("保存に成功した場合は、errorを返さないこと", func(t *testing.T) {
		repo := MockSuccessCaseFavoriteRepository{}
		command := TweetFavoriteEventCommandProcessor{
			favoriteRepository: repo,
		}
		user, tweet, err := SetUp()

		if err != nil {
			t.Fatal(err)
		}

		if result := command.Execute(user, tweet); result != nil {
			t.Fatalf("expected not error, but error happen. error: %v", err)
		}
	})

	t.Run("保存に失敗した場合は、errorを返すこと", func(t *testing.T) {
		repo := MockFailedCaseFavoriteRepository{}
		command := TweetFavoriteEventCommandProcessor{
			favoriteRepository: repo,
		}
		user, tweet, err := SetUp()

		if err != nil {
			t.Fatal(err)
		}

		if result := command.Execute(user, tweet); result == nil {
			t.Fatalf("expected error happen, but error is nil")
		}
	})
}
