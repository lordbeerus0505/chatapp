package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	cors "github.com/rs/cors/wrapper/gin"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "abhiram"
	dbname = "chatapp"
)

type LoginInfo struct {
	Email    string
	Password string
}

type RegistrationInfo struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

const AUTH_INSERT = "INSERT INTO auth (f_name, l_name, email, password) VALUES ($1, $2, $3, $4)"

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/login", func(context *gin.Context) {
		body := LoginInfo{}
		// using BindJson method to serialize body with struct
		if err := context.BindJSON(&body); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body.Email, body.Password)
		context.JSON(http.StatusOK, &body)
	})

	r.POST("/register", func(context *gin.Context) {
		body := RegistrationInfo{}

		if err := context.BindJSON(&body); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body.FirstName, body.LastName, body.Email, body.Password)
		if registerAccount(body.FirstName, body.LastName, body.Email, body.Password) {
			fmt.Println("Successfully Registered the account")
			context.JSON(http.StatusOK, "Yes, It was a success")
		} else {
			context.JSON(http.StatusBadRequest, false)
		}

	})
	r.Run()

}

func registerAccount(fname string, lname string, email string, password string) bool {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=password dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close() // Wait till end of function before executing this.
	_, err = db.Exec(AUTH_INSERT, fname, lname, email, password)
	if err != nil {
		panic(err) // Error Handling and sending back to the user is yet to be implemented.
	}

	fmt.Println("Successfully created account!")
	return true
}
func validateLogin(c *gin.Context) {
	fmt.Println("Hitting the backend successfully", c)
}
