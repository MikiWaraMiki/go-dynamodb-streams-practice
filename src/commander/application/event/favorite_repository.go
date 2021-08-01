package application_event

type FavoriteRepository interface {
	Store(event AddTweetFavoriteEvent) error
}
