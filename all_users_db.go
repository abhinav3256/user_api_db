package main

import "log"

func getAllUserFromDB() ([]User, error) {
	users := []User{}
	userSQL := "SELECT id, name, email,user_id, phone, city,  password FROM users"

	rows, err := DB.Query(userSQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return users, err
	}
	defer rows.Close()
	user := User{}
	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Email, &user.UserID, &user.Phone, &user.City, &user.Password)
		users = append(users, user)
	}
	return users, err
}
