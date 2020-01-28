package service

import (
	"github.com/hawltu/project1/entity"
	//"github.com/hawltu/project1/user"
	 "github.com/hawltu/project1/item"
)

// ItemServiceImpl implements menu.ItemService interface
type ItemServiceImpl struct {
	ItemRepo item.ItemRepository
}


//
// NewItemServiceImpl will create new ItemService object
func NewItemServiceImpl(CatRepo item.ItemRepository) item.ItemService {
	return &ItemServiceImpl{ItemRepo: CatRepo}
}

// Items returns list of Items
func (cs *ItemServiceImpl) Items() ([]entity.Item, error) {

	Items, err := cs.ItemRepo.Items()

	if err != nil {
		return nil, err
	}
	return Items, nil
}

// StoreItem persists new Item information
func (cs *ItemServiceImpl) StoreItem(item1 *entity.Item) (*entity.Item,[]error) {

	itms,err := cs.ItemRepo.StoreItem(item1)

	if len(err) > 0 {
		return nil,err
	}
	return itms,nil
}

// Item returns a Item object with a given id
func (cs *ItemServiceImpl) Item(id int) (entity.Item, error) {

	c, err := cs.ItemRepo.Item(id)

	if err != nil {
		return c, err
	}
	return c, nil
}

// UpdateItem updates a cateogory with new data
func (cs *ItemServiceImpl) UpdateItem(Item entity.Item) error {

	err := cs.ItemRepo.UpdateItem(Item)

	if err != nil {
		return err
	}

	return nil
}

// DeleteItem delete a Item by its id
func (cs *ItemServiceImpl) DeleteItem(id int) error {

	err := cs.ItemRepo.DeleteItem(id)
	if err != nil {
		return err
	}
	return nil
}
