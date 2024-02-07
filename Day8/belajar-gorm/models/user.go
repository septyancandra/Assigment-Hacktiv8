package models

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null; unique; type:varchar(191)"`
	Product   []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null; unique; type:varchar(191)"`
	Brand     string `gorm:"not null; unique; type:varchar(191)"`
	UserId    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Product Before Create()")

	if len(p.Name) < 4 {
		err = errors.New("Product Name is too short")
	}
	return
}
