package repository


import (
	"database/sql"
	"errors"

	"github.com/hawltu/project1/entity"
)

type ItemRepositoryImpl struct {
	conn *sql.DB
}

func NewItemRepositoryImpl(Conn *sql.DB) *ItemRepositoryImpl {
	return &ItemRepositoryImpl{conn: Conn}
}

func (cri *ItemRepositoryImpl) Items() ([]entity.Item, error) {

	rows, err := cri.conn.Query("SELECT * FROM  Items;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.Item{}

	for rows.Next() {
		item := entity.Item{}
		err = rows.Scan(&item.id,  &item.name , &item.catagory, &item.subcatagory, &item.prie, &item.quanity)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, item)
	}
	

		return ctgs, nil
}

func (cri *ItemRepositoryImpl) Item(id int) (entity.Item, error) {

		row := cri.conn.QueryRow("SELECT * FROM Items WHERE id = $1", id)

		item := &entity.Item{}

	err := row.Scan(&item.id,&item.name,&item.catagory,&item.subcatagory,&item.price,&item.quanity)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (cri *ItemRepositoryImpl) UpdateItem(c entity.Item) error {
	//
	_, err := cri.conn.Exec("UPDATE Items SET name=$1,price=$2, catagory=$3 subcatagory = $4 quantity = $5 WHERE id=$6", c.name, c.price, c.catagory, c.subcatagory, c.quantity, c.id)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteCategory removes a category from a database by its id
func (cri *ItemRepositoryImpl) DeleteItem(id int) error {

	_, err := cri.conn.Exec("DELETE FROM Items WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// StoreCategory stores new category information to database
func (cri *ItemRepositoryImpl) StoreItem(c entity.Item) error {

	_, err := cri.conn.Exec("INSERT INTO Items (name ,price,quantity, catagory,subcatagory) values($1, $2, $3,$4,$5)", c.name, c.price, c.quantity, c.catagory, c.subcatagory)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil

}
