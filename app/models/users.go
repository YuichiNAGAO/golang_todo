package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.Password), time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid, name, email, password, created_at from users where id = ?`

	var created_at []uint8

	err = Db.QueryRow(cmd, id).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &created_at)

	created_at_str := string(created_at)                                      // []uint8 を string に変換する
	created_at_time, err := time.Parse("2006-01-02 15:04:05", created_at_str) // string を time.Time に変換する
	if err != nil {
		log.Fatalln(err)
	}

	user.CreatedAt = created_at_time // 変換した created_at の値を user.CreatedAt に代入する

	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`

	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
