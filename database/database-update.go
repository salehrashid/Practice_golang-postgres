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

	if err := updateData(db); err != nil {
		return
	}
}

func updateData(db *sql.DB) error {
	sqlStatement := `UPDATE users SET name = $1, email = $2 WHERE id = $3`

	newName := "rizky"
	newEmail := "admin@admin.com"
	userID := 8

	_, err := db.Exec(sqlStatement, newName, newEmail, userID)
	if err != nil {
		panic(err)
	}

	println("Data updated successfully.")

	return nil
}
