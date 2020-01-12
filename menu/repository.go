package menu

import "github.com/hawltu/project1/entity"

<<<<<<< HEAD

type ItemRepository interface {
	Items() ([]entity.Item, error)
	Item(id int) (entity.Item, error)
	UpdateItem(item entity.Item) error
	DeleteItem(id int) error
	StoreItem(item entity.Item) error
}
type UserRepository interface {
	Users()([]entity.User, error)
	User(id int) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
	StoreUser(user *entity.User) error
}





=======
// CategoryRepository specifies menu category related database operations
type CategoryRepository interface {
	Users() ([]entity.User, error)
	User(id int) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id int) error
	StoreUser(user entity.User) error
}
>>>>>>> c88e25e2abc34e1c0b678c5686509ebca8fed30c
