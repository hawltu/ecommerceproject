package menu

import "github.com/hawltu/project1/entity"

// CategoryService specifies food menu category services
type UserService interface {
	Users() ([]entity.User, error)
	User(id int) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
	StoreUser(user *entity.User) error
}


type ItemService interface {
	Items() ([]entity.Item, error)
	Item(id int) (entity.Item, error)
	UpdateItem(user entity.Item) error
	DeleteItem(id int) error
	StoreItem(item entity.Item) error
}



