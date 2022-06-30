package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	DB *sql.DB
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	UserID   string `json:"user_id"`
	City     string `json:"city"`
	Password string `json:"password" binding:"required"`
}

var Data = make(map[string]User)

func main() {

	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRoutes(r *gin.Engine) {
	r.GET("/user/:user_id", GetUserByUserID)
	r.GET("/user", GetAllUser)
	r.POST("/user", CreateUser)
	//r.PUT("/user/:user_id", MyMiddleware(), UpdateUser)
	r.DELETE("/user/:user_id", deleteUser)
}

//GetUserByUserID function
func GetUserByUserID(c *gin.Context) {
	//records := readCsvFile("./movies.csv")
	userID, ok := c.Params.Get("user_id")
	if ok == false {
		res := gin.H{
			"error": "user_id is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	//	var user User

	//user = getUserByID(userID)
	user, _ := getUserByIDFromDB(userID)
	//user, _ := create_user_db()
	res := gin.H{
		"user":    user,
		"user_id": userID,
	}
	c.JSON(http.StatusOK, res)
}

//GetAllUser function
func GetAllUser(c *gin.Context) {
	res := gin.H{
		"user": Data,
	}
	c.JSON(http.StatusOK, res)
}

//CreateUser POST
func CreateUser(c *gin.Context) {
	reqBody := User{}
	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if len(reqBody.Phone) != 13 {
		res := gin.H{
			"error": "phone number must be 13 digit",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	// if isUnique(reqBody.UserID) {
	// 	res := gin.H{
	// 		"error": "UserId already Exist",
	// 	}
	// 	c.JSON(http.StatusBadRequest, res)
	// 	return
	// }

	// if UniqueEmail(reqBody.Email) {
	// 	res := gin.H{
	// 		"error": "Email id already Exist",
	// 	}
	// 	c.JSON(http.StatusBadRequest, res)
	// 	return
	// }
	//Data[reqBody.UserID] = reqBody
	result, _ := create_user_db(reqBody)

	res := gin.H{
		"success": true,
		"result":  result,
	}
	c.JSON(http.StatusOK, res)
	return
}

//Update User PUT
func UpdateUser(c *gin.Context) {

	userID, ok := c.Params.Get("user_id")
	if ok == false {
		res := gin.H{
			"error": "user_id is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}

	reqBody := User{}

	err := c.Bind(&reqBody)
	if err != nil {
		res := gin.H{
			"error": "user id is required",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if reqBody.UserID != userID {
		res := gin.H{
			"error": "Invalid UserID",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//password := c.GetHeader("password")
	// if !checkPassword(userID, password) {
	// 	res := gin.H{
	// 		"success": false,
	// 		"message": "Incorrect password",
	// 	}
	// 	c.JSON(http.StatusBadRequest, res)
	// 	return
	// }

	Data[userID] = reqBody
	res := gin.H{
		"success": true,
		"user":    reqBody,
	}
	c.JSON(http.StatusOK, res)
	return
}

//Update User PUT
func deleteUser(c *gin.Context) {
	userID, ok := c.Params.Get("user_id")
	fmt.Println(userID)
	fmt.Println(Data)
	if ok == false {
		res := gin.H{
			"error": "user_id is missing",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	password := c.GetHeader("password")
	if !checkPassword(userID, password) {
		res := gin.H{
			"success": false,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result := delete_user(userID)
	res := gin.H{
		"success": true,
		"message": result,
	}
	c.JSON(http.StatusOK, res)
	return
}

func MyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// middleware

		userID, _ := c.Params.Get("user_id")
		password := c.GetHeader("password")
		if _, ok := Data[userID]; ok {

			if password != Data[userID].Password {
				res := gin.H{
					"success": false,
					"message": "Incorrect password",
				}
				c.JSON(http.StatusBadRequest, res)
				c.Abort()
			}
		}
		c.Next()
	}
}
