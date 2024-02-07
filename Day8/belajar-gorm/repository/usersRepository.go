package repository

import (
	"belajar-gorm/database"
	"belajar-gorm/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func CreateUser(email string) {
	db := database.GetDB()

	user := models.User{
		Email: email,
	}
	err := db.Create(&user).Error
	if err != nil {
		fmt.Println("Error creating user data :", err)
		return
	}
	fmt.Println("New User data : ", user)
}

func GetUserById(id uint) {
	db := database.GetDB()
	user := models.User{}
	err := db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User Data Not Found ", err)
			return
		}
		print("Error Finding user: ", err)
	}
	fmt.Printf("User Data : %+v \n ", user)
}

func UpdateUserById(id uint, email string) {
	db := database.GetDB()
	user := models.User{}
	err := db.Model(&user).Where("id = ?", id).Updates(models.User{
		Email: email,
	}).Error

	if err != nil {
		fmt.Println("Error updating user data :", err)
		return
	}
	fmt.Printf("Update User's email: %+v \n ", user.Email)
}
