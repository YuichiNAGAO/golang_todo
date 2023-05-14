package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"github.com/YuichiNAGAO/golang_todo/config"
	"github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

var err error

const (
	tableName = "users"
)

func init() {
	log.Println(config.Config.SQLDriver)
	log.Println(config.Config.DbName)
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		uuid VARCHAR(255) NOT NULL UNIQUE,
		name VARCHAR(255),
		email VARCHAR(255),
		password VARCHAR(255),
		created_at DATETIME)`, tableName)

	fmt.Println(cmd)
	Db.Exec(cmd)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
