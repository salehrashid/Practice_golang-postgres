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

	sqlStatement := `
INSERT INTO users (name, email)
VALUES ($1, $2)
RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, "saleh", "a@a.com").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}
