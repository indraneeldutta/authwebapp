package main

import (
	"database/sql"
	"log"

	"github.com/markbates/goth"
)

type User struct {
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedOn string `json:"created_on"`
	Meta      string `json:"meta"`
}

// AddUserDetails add the user details to the postgres DB
func AddUserDetails(user goth.User, provider string) {
	var exist bool
	checkStatement := `select exists(select 1 from auth.user.auth where email='indraneel.dutta@hotmail.com')`
	err := db.QueryRow(checkStatement).Scan(&exist)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	if exist {
		var meta string
		getMeta := `select meta from auth.user.auth where email=$1`
		err := db.QueryRow(getMeta, user.Email).Scan(&meta)
		if err != nil && err != sql.ErrNoRows {
			log.Fatal(err)
		}

		meta = meta + "," + provider
		update := `update auth.user.auth set meta=$1 where email=$2`
		_, err = db.Exec(update, meta, user.Email)
		if err != nil && err != sql.ErrNoRows {
			log.Fatal(err)
		}
	} else {
		sqlStatement := `
		INSERT INTO auth.user.auth (name, email, phone_number, created_on, meta)
		VALUES ($1, $2, $3, current_timestamp, $4)
		RETURNING user_id`

		id := 0
		err := db.QueryRow(sqlStatement, user.Name, user.Email, "", provider).Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// GetAllUsers will retrieve all users from Database and return as sql rows format
func GetAllUsers() *sql.Rows {
	getUsers := `select * from auth.user.auth`
	rows, err := db.Query(getUsers)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	return rows
}
