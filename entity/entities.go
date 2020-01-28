package entity

import (
	"github.com/jinzhu/gorm"
)
type User struct {
	gorm.Model
	//ID        uint `gorm:"type:integer;not null"`
	LName     string `gorm:"type:varchar(255);not null"`
	FName     string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	UserName  string `gorm:"type:varchar(255);not null"`
	Mobile    string `gorm:"type:varchar(255);not null"`
	Address   string `gorm:"type:varchar(255);not null"`
	Shopname  string `gorm:"type:varchar(255);not null"`
	Image  string `gorm:"type:varchar(255);not null"`
}

// UserSession represents user sessions
type UserSession struct {
	gorm.Model
	//ID         uint 
	UserID     uint
	UUID       string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}
type Item struct {
	gorm.Model
	//ID          int   
	Name        string  `gorm:"type:varchar(255);not null"`
	Catagory    string  `gorm:"type:varchar(255);not null"`
	Subcatagory string  `gorm:"type:varchar(255);not null"`
	Price       int    `gorm:"type:integer;not null"`
	Quantity    int     `gorm:"type:integer;not null"`   
	Image       string  `gorm:"type:varchar(255);not null"`
	UserID      uint 
}
