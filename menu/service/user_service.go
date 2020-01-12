package service

import (
	"github.com/hawltu/project1/entity"
	"github.com/hawltu/project1/menu"
)

<<<<<<< HEAD
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
=======
// CategoryServiceImpl implements menu.CategoryService interface
type UserServiceImpl struct {
	userRepo menu.UserRepository
}

// NewCategoryServiceImpl will create new CategoryService object
func NewUserServiceImpl(CatRepo menu.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{categoryRepo: CatRepo}
}

// Categories returns list of categories
func (cs *UserServiceImpl) User() ([]entity.User, error) {

	categories, err := cs.userRepo.Users()
>>>>>>> c88e25e2abc34e1c0b678c5686509ebca8fed30c

	if err != nil {
		return nil, err
	}

<<<<<<< HEAD
	return Users, nil
}

// StoreUser persists new User information
func (cs *UserserviceImpl) StoreUser(User *entity.User) error {

	err := cs.UserRepo.StoreUser(User)
=======
	return categories, nil
}

// StoreCategory persists new category information
func (cs *CategoryServiceImpl) StoreUser(category entity.User) error {

	err := cs.categoryRepo.UserCategory(category)
>>>>>>> c88e25e2abc34e1c0b678c5686509ebca8fed30c

	if err != nil {
		return err
	}

	return nil
}

<<<<<<< HEAD
// User returns a User object with a given id
func (cs *UserserviceImpl) User(id int) (entity.User, error) {

	c, err := cs.UserRepo.User(id)
=======
// Category returns a category object with a given id
func (cs *UserServiceImpl) User(id int) (entity.User, error) {

	c, err := cs.userRepo.User(id)
>>>>>>> c88e25e2abc34e1c0b678c5686509ebca8fed30c

	if err != nil {
		return c, err
	}

	return c, nil
}

<<<<<<< HEAD
// UpdateUser updates a cateogory with new data
func (cs *UserserviceImpl) UpdateUser(User entity.User) error {

	err := cs.UserRepo.UpdateUser(User)
=======
// UpdateCategory updates a cateogory with new data
func (cs *CategoryServiceImpl) UpdateCategory(category entity.Category) error {

	err := cs.categoryRepo.UpdateCategory(category)
>>>>>>> c88e25e2abc34e1c0b678c5686509ebca8fed30c

	if err != nil {
		return err
	}

	return nil
}

<<<<<<< HEAD
// DeleteUser delete a User by its id
func (cs *UserserviceImpl) DeleteUser(id int) error {

	err := cs.UserRepo.DeleteUser(id)
=======
// DeleteCategory delete a category by its id
func (cs *CategoryServiceImpl) DeleteCategory(id int) error {

	err := cs.categoryRepo.DeleteCategory(id)
>>>>>>> c88e25e2abc34e1c0b678c5686509ebca8fed30c
	if err != nil {
		return err
	}
	return nil
}
