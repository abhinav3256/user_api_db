package main

import (
	"database/sql"
	"log"
)

func getUserByID(userID string) User {
	user := User{}
	if user, ok := Data[userID]; ok {
		return user
	}

	return user
}

func getUserByIDFromDB(userID string) (User, error) {
	var err error
	DB, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	//defer DB.Close()
	user := User{}
	userSQL := "SELECT id, name, userid, phone, city,  password FROM users WHERE userid = $1"

	err = DB.QueryRow(userSQL, userID).Scan(&user.ID, &user.Name, &user.UserID, &user.Phone, &user.City, &user.Password)
	if err != nil {
		log.Println("Failed to execute query: ", err)
	}
	return user, err
}
