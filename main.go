package main

import (
	"fmt"
	"log"

	"github.com/YuichiNAGAO/golang_todo/app/models"
)

func main() {
	// fmt.Println(config.Config)

	// log.Println("test")

	fmt.Println(models.Db)

	// u := &models.User{}

	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	u, err := models.GetUser(1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(u)

	u.Name = "test2"
	u.Email = "test2@example.com"
	u.UpdateUser()
	u, err = models.GetUser(1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(u)

	// u.DeleteUser()

}
