package main

import (
	"fmt"

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

	// u, err := models.GetUser(2)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(u)

	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, err = models.GetUser(2)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(u)

	// u.DeleteUser()

	// u, _ := models.GetUser(2)
	// u.CreateTodo("first todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	user2, _ := models.GetUser(2)
	todos, _ := user2.GetTodosByUser()
	for _, v := range todos {
		fmt.Println(v)
	}

}
