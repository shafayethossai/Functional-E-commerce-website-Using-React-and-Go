package user

import (
	"first-program/domain"
)

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service { // injecting user repository
	return &service{
		usrRepo: usrRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) { // call the Create method of the user repository to create a new user in the database
	usr, err := svc.usrRepo.Create(user)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}
	return usr, nil
}
func (svc *service) Find(email string, pass string) (*domain.User, error) { // call the Find method of the user repository to find a user in the database by email and password
	usr, err := svc.usrRepo.Find(email, pass)
	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, nil
	}
	return usr, nil
}
