package domain_model_user

type UserRepository interface {
	FindById(id *UserID) (*User, error)
}
