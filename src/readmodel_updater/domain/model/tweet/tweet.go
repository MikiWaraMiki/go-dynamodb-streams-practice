package domain_model_tweet

type Tweet struct {
	id      *TweetID
	content string
}

func NewTweet(id *TweetID, content string) *Tweet {
	return &Tweet{
		id:      id,
		content: content,
	}
}

func (t Tweet) Id() string {
	return t.id.value
}

func (t Tweet) Content() string {
	return t.content
}
