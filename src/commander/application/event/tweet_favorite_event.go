package application_event

import "time"

type TweetFavoriteEventBody struct {
	tweetId string
}

type TweetFavoriteEvent struct {
	userId         string
	sequenceNumber int
	eventType      string
	body           TweetFavoriteEventBody
	createdAt      time.Time
}

type TweetEventProvider struct {
	userId  string
	version string
}
