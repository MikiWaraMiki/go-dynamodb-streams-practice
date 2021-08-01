package domain_model_favorite

type FavoriteRepository interface {
	InsertFavorite(favorite *Favorite) error
}
