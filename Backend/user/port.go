package user

import (
	"first-program/domain"
	userHandler "first-program/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}
type UserRepo interface { // signatures of the methods that the user repository should implement
	Create(user domain.User) (*domain.User, error)
	Find(email string, pass string) (*domain.User, error)
}
