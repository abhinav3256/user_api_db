package main

import (
	"database/sql"
	"fmt"
	"log"
)

func delete_user_db(userid string) (bool, bool) {

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

	var result = true
	//sqlStatement := `
	//INSERT INTO users(id, name, email, phone, userid, city, password)
	//VALUES ($1, $2, $3, $4,$5,$6,$7)`
	sqlStatement := `DELETE FROM users WHERE id = $1`
	user, err2 := DB.Exec(sqlStatement, userid)

	if err2 != nil {
		log.Fatal("ERror in DELETE: ", err2)
	}
	fmt.Println(user)
	return result, err_responce
}
