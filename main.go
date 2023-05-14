package main

import (
	"fmt"

	"github.com/YuichiNAGAO/golang_todo/app/models"
)

func main() {
	// fmt.Println(config.Config)

	// log.Println("test")

	fmt.Println(models.Db)

	u := &models.User{}

	u.Name = "test"
	u.Email = "test@example.com"
	u.Password = "testtest"
	fmt.Println(u)

	u.CreateUser()
}
