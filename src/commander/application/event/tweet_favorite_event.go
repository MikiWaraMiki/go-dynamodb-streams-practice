package application_event

import "time"

type TweetFavoriteEventBody struct {
	userId  string
	tweetId string
}

type AddTweetFavoriteEvent struct {
	UserId         string
	SequenceNumber int
	EventType      string
	Body           TweetFavoriteEventBody
	CreatedAt      time.Time
}
