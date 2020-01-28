package service

import (
	"github.com/hawltu/project1/user"
	"github.com/hawltu/project1/entity"
)

// UserService implements model.UserRepository interface
type UserService struct {
	userRepo user.UserRepository
}

// NewUserService will create new UserService object
func NewUserService(UserRepos user.UserRepository) user.UserService {
	return &UserService{userRepo: UserRepos}
}

// Users returns list of users
func (us *UserService) Users() ([]entity.User, []error) {

	users, errs := us.userRepo.Users()

	if len(errs) > 0 {
		return nil, errs
	}

	return users, nil
}

// StoreUser persists new user information
func (us *UserService) StoreUser(user *entity.User) (*entity.User, []error) {

	usr, errs := us.userRepo.StoreUser(user)

	if len(errs) > 0 {
		return nil,errs
	}

	return usr, nil
}
func (us *UserService) UserByUserName(username string) (*entity.User, []error) {
	usr, errs := us.userRepo.UserByUserName(username)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// User returns a user object with a given id
func (us *UserService) User(id uint) (*entity.User, []error) {

	user, errs := us.userRepo.User(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return user, nil
}

// UpdateUser updates a user with new data
func (us *UserService) UpdateUser(user *entity.User) (*entity.User, []error) {

	usr, errs := us.userRepo.UpdateUser(user)

	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

// DeleteUser deletes a user by its id
func (us *UserService) DeleteUser(id uint) (*entity.User, []error) {

	usr, errs := us.userRepo.DeleteUser(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return usr, nil
}

// StoreSession persists new session information
func (us *UserService) StoreSession(session *entity.UserSession) (*entity.UserSession, []error) {

	s, errs := us.userRepo.StoreSession(session)

	if len(errs) > 0 {
		return nil, errs
	}

	return s, nil
}

// DeleteSession delete a session by its id
func (us *UserService) DeleteSession(uuid string) (*entity.UserSession, []error) {

	s, errs := us.userRepo.DeleteSession(uuid)

	if len(errs) > 0 {
		return nil, errs
	}
	return s, nil
}

// Session returns a session object with a given id
func (us *UserService) Session(uuid string) (*entity.UserSession, []error) {

	s, errs := us.userRepo.Session(uuid)

	if len(errs) > 0 {
		return nil, errs
	}

	return s, nil
}
