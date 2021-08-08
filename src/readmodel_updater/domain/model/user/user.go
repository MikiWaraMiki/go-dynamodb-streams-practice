package domain_model_user

type User struct {
	id *UserID
}

func NewUser(id *UserID) *User {
	return &User{
		id: id,
	}
}

func (u *User) Id() string {
	return u.id.Value()
}
