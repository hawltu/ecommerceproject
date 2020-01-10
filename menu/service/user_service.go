package service

import (
	"github.com/hawltu/project1/entity"
	"github.com/hawltu/project1/menu"
)

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

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// StoreCategory persists new category information
func (cs *CategoryServiceImpl) StoreUser(category entity.User) error {

	err := cs.categoryRepo.UserCategory(category)

	if err != nil {
		return err
	}

	return nil
}

// Category returns a category object with a given id
func (cs *UserServiceImpl) User(id int) (entity.User, error) {

	c, err := cs.userRepo.User(id)

	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateCategory updates a cateogory with new data
func (cs *CategoryServiceImpl) UpdateCategory(category entity.Category) error {

	err := cs.categoryRepo.UpdateCategory(category)

	if err != nil {
		return err
	}

	return nil
}

// DeleteCategory delete a category by its id
func (cs *CategoryServiceImpl) DeleteCategory(id int) error {

	err := cs.categoryRepo.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}
