package service

import (
	"github.com/hawltu/project1/entity"
	"github.com/hawltu/project1/menu"
)

type UserserviceImpl struct {
	UserRepo menu.UserRepository
}

// NewUserserviceImpl will create new Userservice object
func NewUserserviceImpl(CatRepo menu.UserRepository) *UserserviceImpl {
	return &UserserviceImpl{UserRepo: CatRepo}
}

// Users returns list of Users
func (cs *UserserviceImpl) Users() ([]entity.User, error) {

	Users, err := cs.UserRepo.Users()

	if err != nil {
		return nil, err
	}

	return Users, nil
}

// StoreUser persists new User information
func (cs *UserserviceImpl) StoreUser(User *entity.User) error {

	err := cs.UserRepo.StoreUser(User)

	if err != nil {
		return err
	}

	return nil
}

// User returns a User object with a given id
func (cs *UserserviceImpl) User(id int) (entity.User, error) {

	c, err := cs.UserRepo.User(id)

	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateUser updates a cateogory with new data
func (cs *UserserviceImpl) UpdateUser(User entity.User) error {

	err := cs.UserRepo.UpdateUser(User)

	if err != nil {
		return err
	}

	return nil
}

// DeleteUser delete a User by its id
func (cs *UserserviceImpl) DeleteUser(id int) error {

	err := cs.UserRepo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
