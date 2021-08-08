package domain_model_favorite_tweet

type FavoriteTweetRepository interface {
	InsertOne(favoriteTweet *FavoriteTweet) error
}
