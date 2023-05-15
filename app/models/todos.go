package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateTodo(content string) (err error) {
	cmd := `insert into todos (
		content,
		user_id,
		created_at) values (?, ?, ?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func GetTodo(id int) (todo Todo, err error) {
	cmd := `select id, content, user_id, created_at from todos where id = ?`
	todo = Todo{}

	var created_at []uint8
	err = Db.QueryRow(cmd, id).Scan(&todo.ID, &todo.Content, &todo.UserID, &created_at)

	created_at_str := string(created_at)                                      // []uint8 を string に変換する
	created_at_time, err := time.Parse("2006-01-02 15:04:05", created_at_str) // string を time.Time に変換する
	if err != nil {
		log.Fatalln(err)
	}

	todo.CreatedAt = created_at_time

	return todo, err
}

func GetTodos() (todos []Todo, err error) {
	cmd := `select id, content, user_id from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}

	rows.Close()

	return todos, err
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, content, user_id from todos where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Content, &todo.UserID)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}

	return todos, err
}
