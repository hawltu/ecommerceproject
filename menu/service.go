package menu

import "github.com/hawltu/project1/entity"

// CategoryService specifies food menu category services
<<<<<<< HEAD
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



=======
type CategoryService interface {
	users() ([]entity.User, error)
	User(id int) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
	StoreUser(user entity.User) error
}
>>>>>>> c88e25e2abc34e1c0b678c5686509ebca8fed30c
