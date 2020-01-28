package repository

import (
	"errors"

	"github.com/hawltu/project1/entity"
	"github.com/hawltu/project1/user"
	"github.com/jinzhu/gorm"
)

// MockUserRepo implements the menu.UserRepository interface
type MockUserRepo struct {
	conn *gorm.DB
}

// NewMockUserRepo will create a new object of MockUserRepo
func NewMockUserRepo(db *gorm.DB) user.UserRepository {
	return &MockUserRepo{conn: db}
}

// Users returns all fake categories
func (mCatRepo *MockUserRepo) Users() ([]entity.User, []error) {
	ctgs := []entity.User{entity.UserMock}
	return ctgs, nil
}

// User retrieve a fake user with id 1
func (mCatRepo *MockUserRepo) User(id uint) (*entity.User, []error) {
	ctg := entity.UserMock
	if id == 1 {
		return &ctg, nil
	}
	return nil, []error{errors.New("Not found")}
}

// UpdateUser updates a given fake user
func (mCatRepo *MockUserRepo) UpdateUser(user *entity.User) (*entity.User, []error) {
	cat := entity.UserMock
	return &cat, nil
}

// DeleteUser deletes a given user from the database
func (mCatRepo *MockUserRepo) DeleteUser(id uint) (*entity.User, []error) {
	cat := entity.UserMock
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &cat, nil
}

// StoreUser stores a given mock user
func (mCatRepo *MockUserRepo) StoreUser(user *entity.User) (*entity.User, []error) {
	cat := user
	return cat, nil
}
func (mCatRepo *MockUserRepo) DeleteSession(uuid string) (*entity.UserSession, []error) {
	s, errs := mCatRepo.DeleteSession(uuid)
	if (len(errs) >0){
		return nil,nil
	}
	return s, nil
}
func (userRepo *MockUserRepo) UserByUserName(username string) (*entity.User, []error) {
	user := entity.User{}
	errs := userRepo.conn.Find(&user, "user_name=?", username).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &user, errs
}
// StoreSession stores a given session in the database
func (uRepo *MockUserRepo) StoreSession(session *entity.UserSession) (*entity.UserSession, []error) {
	s := session
	errs := uRepo.conn.Create(s).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return s, errs
}
func (uRepo *MockUserRepo)Session(uuid string) (*entity.UserSession, []error){
     return nil,nil
}
