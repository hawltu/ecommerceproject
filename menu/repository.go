package menu

import "github.com/hawltu/project1/entity"


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





