package domain_model_tweet

import "errors"

type TweetID struct {
	value string
}

func NewTweetId(id string) (*TweetID, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	return &TweetID{
		value: id,
	}, nil
}

func (tid *TweetID) Value() string {
	return tid.value
}
