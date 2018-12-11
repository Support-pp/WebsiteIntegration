package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var conString = os.Getenv("DBCON")

func checkIfEmailExist(email string) bool {
	db, err := sql.Open("mysql", conString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("")
	results, err := db.Query("SELECT * FROM token WHERE email=?", email)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag APIToken
		err = results.Scan(&tag.UID, &tag.Email, &tag.Token, &tag.blocked)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("User found with same email :: ", tag.UID)
		db.Close()
		return true
	}

	db.Close()
	return false
}

func createNewAPIToken(email string, token string) {
	db, err := sql.Open("mysql", conString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO token (email, token) VALUES (?, ?)", email, token)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	insert.Close()
	db.Close()
	return
}
