package favorite_tweet

type FavoriteTweet struct {
	user_id  string
	tweet_id int
	body     string
}

func NewFavoriteTweet(user_id *UserID, tweet *Tweet) *FavoriteTweet {
	return *FavoriteTweet{
		user_id:  user_id.Value(),
		tweet_id: tweet.Id(),
		body:     tweet.Body(),
	}
}
