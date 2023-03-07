package models

import (
	"app/config"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

var DB *sql.DB
var err error

const (
	tableUser = "users"
)

func init() {
	DB, err = sql.Open(config.Config.SQLDriver, config.Config.DbUser+":"+config.Config.DbPass+"@tcp(mysql:3306)/"+config.Config.DbName+"?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	//defer func(DB *sql.DB) {
	//	err = DB.Close()
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//}(DB)

	err = DB.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
 id INTEGER PRIMARY KEY AUTO_INCREMENT,
 uuid VARCHAR(255) NOT NULL UNIQUE,
 name VARCHAR(255),
 email VARCHAR(255),
 password VARCHAR(255),
 created_at DATETIME
 )`, tableUser)

	_,err = DB.Exec(cmdU)
	if err != nil {
		log.Fatalln(err)
	}
}

func CreateUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
