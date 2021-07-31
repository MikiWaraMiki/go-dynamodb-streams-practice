package domain_model_tweet

type Tweet struct {
	id *TweetId
}

func NewTweet(id *TweetId) *Tweet {
	return &Tweet{
		id: id,
	}
}

func (t Tweet) Id() string {
	return t.id.value
}
