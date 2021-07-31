package domain_model_tweet

import "errors"

type TweetId struct {
	value string
}

func NewTweetId(id string) (*TweetId, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	return &TweetId{
		value: id,
	}, nil
}
