package repository

import (
	"belajar-gorm/database"
	"belajar-gorm/models"
	"fmt"
)

func CreateProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	product := models.Product{
		UserId: userId,
		Brand:  brand,
		Name:   name,
	}
	err := db.Create(&product).Error
	if err != nil {
		fmt.Println("Error Creating product Data: ", err.Error())
		return
	}
	fmt.Println("New Product Data ", product)
}

func GetUsersWithProducts() {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Product").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user datas with products: ", err.Error())
		return
	}
	fmt.Println("user datas with product ")
	fmt.Printf("%+v", users)
}

func DeleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}
	err := db.Where("id = ?", id).Delete(&product).Error
	if err != nil {
		fmt.Println("Error Deleting product :", err.Error())
		return
	}
	fmt.Println("Product with id %d has been successfuly deleted ", id)
}
