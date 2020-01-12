package menu

import "github.com/hawltu/project1/entity"

// CategoryService specifies food menu category services
type UserService interface {
	users() ([]entity.User, error)
	User(id int) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
	StoreUser(user entity.User) error
}

type ItemService interface {
	Itemss() ([]entity.User, error)
	Item(id int) (entity.User, error)
	UpdateItem(user entity.User) error
	DeleteItem(id int) error
	StoreItem(user entity.User) error
}

