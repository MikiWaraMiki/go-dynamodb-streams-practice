package application_event

import "time"

type TweetFavoriteEventBody struct {
	userId  string
	tweetId string
}

type AddTweetFavoriteEvent struct {
	userId         string
	sequenceNumber int
	eventType      string
	body           TweetFavoriteEventBody
	createdAt      time.Time
}
