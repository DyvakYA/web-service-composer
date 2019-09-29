package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// entity User
type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
}

//get user from db
func FindUser(db *sql.DB, id string) (User) {
	row := db.QueryRow("SELECT id, first_name FROM users WHERE id=?", id)

	var user User
	switch err := row.Scan(&user.Id, &user.FirstName); err {
	case sql.ErrNoRows:
		fmt.Println("User not found!")
	case nil:
		fmt.Println("null")
	default:
		panic(err)
	}
	return user
}
