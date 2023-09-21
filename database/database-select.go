package main

/**
Note the underscore _ before the "github.com/lib/pq" import.
This is required to register the PostgreSQL driver with the
database/sql package, even though you won't directly use it
in your code.
*/

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

	if err := selectData(db); err != nil {
		panic(err)
	}
}

func selectData(db *sql.DB) error {
	//execute query
	rows, err := db.Query("SELECT name, email FROM users")
	if err != nil {
		panic(err)
	}

	//retrieve and iterate data
	for rows.Next() {
		var name string
		var email string
		if err := rows.Scan(&name, &email); err != nil {
			panic(err)
		}
		fmt.Printf("Name: %s, Email %s\n", name, email)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	return nil
}
