package menu

import "github.com/hawltu/project1/entity"

// CategoryRepository specifies menu category related database operations
type UserRepository interface {
	Users() ([]entity.User, error)
	User(id int) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
	StoreUser(user entity.User) error
}

type ItemRepository interface {
	Itemss() ([]entity.User, error)
	Item(id int) (entity.User, error)
	UpdateItem(user entity.User) error
	DeleteItem(id int) error
	StoreItem(user entity.User) error
}
