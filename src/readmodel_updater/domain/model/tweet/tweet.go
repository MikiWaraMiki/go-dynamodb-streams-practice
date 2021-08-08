package domain_model_tweet

type Tweet struct {
	id   int
	body string
}

func NewTweet(id int, body string) *Tweet {
	return &Tweet{
		id:   id,
		body: body,
	}
}

func (t Tweet) Id() int {
	return t.id
}

func (t Tweet) Body() string {
	return t.body
}
