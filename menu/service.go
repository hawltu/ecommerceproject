package menu

import "github.com/hawltu/project1/entity"

// CategoryService specifies food menu category services
type CategoryService interface {
	users() ([]entity.User, error)
	User(id int) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
	StoreUser(user entity.User) error
}
