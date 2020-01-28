package repository

import (
	"github.com/hawltu/project1/entity"
	"github.com/hawltu/project1/item"
	"github.com/jinzhu/gorm"
)


type ItemGormRepo struct {
	conn *gorm.DB
}


func NewItemGormRepo(db *gorm.DB) item.ItemRepository{
	return &ItemGormRepo{conn: db}
}

func (pRepo *ItemGormRepo) Items() ([]entity.Item, error) {
	posts := []entity.Item{}
	errs := pRepo.conn.Find(&posts).GetErrors()
	if len(errs) > 0 {
		return nil, errs[1]
	}

	return posts, nil
}

func (pRepo *ItemGormRepo) Item(id int) (entity.Item, error) {
	post := entity.Item{}
	errs := pRepo.conn.First(&post, id).GetErrors()
	if len(errs) > 0 {
		return post, errs[1]
	}
	return post, nil
}

func (pRepo *ItemGormRepo) UpdateItem(post entity.Item) (error) {
	pst := post
	errs := pRepo.conn.Save(pst).GetErrors()
	if len(errs) > 0 {
		return errs[1]
	}
	return nil
}

func (pRepo *ItemGormRepo) DeleteItem(id int) (error) {
	post, errs := pRepo.Item(id)
	if errs != nil{
		return nil
	}
	idd := post.ID
	err := pRepo.conn.Delete(post, idd).GetErrors()
	if err != nil {
		return nil
	}
	return nil
}

// StorePost stores a given post in the database
func (pRepo *ItemGormRepo) StoreItem(post *entity.Item) (*entity.Item,[]error) {
	pst := post
	errs := pRepo.conn.Create(pst).GetErrors()
	if len(errs) > 0{
		return nil,errs
	}
	return pst,errs
}
