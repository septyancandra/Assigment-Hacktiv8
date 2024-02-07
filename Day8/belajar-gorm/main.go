package main

import (
	"belajar-gorm/database"
	"belajar-gorm/repository"
)

func main() {
	database.StartDB()
	//repository.CreateUser("septyancandra28@gmail.com")
	//repository.GetUserById(1)
	//repository.UpdateUserById(1, "coba@gmail.com")
	//repository.CreateProduct(1, "hohoh", "hahah")
	repository.GetUsersWithProducts()
	//repository.DeleteProductById(5)
}
