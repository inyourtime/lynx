package domain

import "lynx/model"

type UserRepository interface {
	Create(user *model.User) (err error)
	Fetch() (users []model.User, err error)
	GetByEmail(email string) (user model.User, err error)
	GetByID(id string) (user model.User, err error)
}
