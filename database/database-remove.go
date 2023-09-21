package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"go-postgres/constants"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		constants.Host, constants.Port, constants.User, constants.Password, constants.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := deleteData(db); err != nil {
		panic(err)
	}
}

func deleteData(db *sql.DB) error {
	sqlStatement := `DELETE FROM users WHERE id = $1`

	_, err := db.Exec(sqlStatement, 5)
	if err != nil {
		return err
	}

	println("Delete data succeeded")

	return nil
}
