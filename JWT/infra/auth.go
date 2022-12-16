package infra

import (
	"database/sql"
	"log"

	"github.com/seipan/Go-Authentic/tree/main/JWT/utils"
)

func CreateUser(db *sql.DB, name string, password string) (int64, error) {
	statement := "INSERT INTO users (name, password) VALUES(?,?)"

	stmt, err := db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer stmt.Close()
	hashPass, err := utils.PasswordEncrypt(password)
	password = hashPass
	res, err := db.Exec(name, password)

	id, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}
	return id, nil
}
