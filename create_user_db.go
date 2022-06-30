package main

import (
	"database/sql"
	"fmt"
	"log"
)

func create_user_db(reqBody User) (string, bool) {

	var err_responce bool
	err_responce = true
	var err error
	DB, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	//defer DB.Close()
	//user := User{}
	//	userSQL := "INSERT INTO users (id, name, email, phone, userid, city, password)
	//	VALUES (2, 'NITISH', 'nitish@gmail.com', '+918452369742', 'nitish', 'bhopal', 'nitish')"

	var u string
	sqlStatement := `
INSERT INTO users(id, name, email, phone, userid, city, password)
VALUES ($1, $2, $3, $4,$5,$6,$7)`
	user, err := DB.Exec(sqlStatement, 3, reqBody.Name, reqBody.Email, reqBody.Phone, reqBody.UserID, reqBody.City, reqBody.Password)

	if err != nil {
		err_responce = false
		return u, err_responce
	}
	fmt.Println(user)
	return u, err_responce
}
