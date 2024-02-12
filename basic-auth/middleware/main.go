package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/student", ActionStudent)
	server := new(http.Server)
	server.Addr = ":8081"

	fmt.Println("server started at localhost:8081")
	server.ListenAndServe()
}

func init() {
	students = append(students, &Student{
		Id:    "s001",
		Name:  "bourne",
		Grade: 2,
	})
}
