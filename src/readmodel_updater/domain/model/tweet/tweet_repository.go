package domain_model_tweet

type TweetRepository interface {
	FindById(id *TweetID) (*Tweet, error)
}
