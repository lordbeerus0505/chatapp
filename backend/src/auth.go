package src

import (
	"database/sql"
	"fmt"
	"net/http"

	"crypto/sha256"
	"encoding/base64"

	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "abhiram"
	dbname = "chatapp"
)

const AUTH_INSERT = "INSERT INTO auth (f_name, l_name, email, password) VALUES ($1, $2, $3, $4)"
const VALID_ACC = "SELECT f_name, l_name, email FROM auth WHERE email=$1 AND password=$2;"
const ACC_EXISTS = "SELECT COUNT(*) FROM auth WHERE email=$1"
const CONTACTS_TABLE = "contacts"

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
type ContactCard struct {
	Email     string
	FirstName string
	LastName  string
}

func AddContact(r *gin.Engine) {
	r.POST("/add-contact", func(context *gin.Context) {
		body := ContactCard{}
		if err := context.BindJSON(&body); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body)
		if userExists(body.Email) {
			// Add to the user's contacts
			addContactDynamo(body)
			context.JSON(http.StatusOK, true)
		} else {
			context.JSON(http.StatusBadRequest, false)
		}
	})
}

func Login(r *gin.Engine) {
	r.POST("/login", func(context *gin.Context) {
		body := LoginInfo{}
		// using BindJson method to serialize body with struct
		if err := context.BindJSON(&body); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body.Email, body.Password)
		card, status := validateLogin(body)

		if status {
			fmt.Println("Success")
			context.JSON(http.StatusOK, &card)
		} else {
			context.JSON(http.StatusBadRequest, false)
		}

	})
}

func Register(r *gin.Engine) {
	r.POST("/register", func(context *gin.Context) {
		body := RegistrationInfo{}

		if err := context.BindJSON(&body); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body.FirstName, body.LastName, body.Email, body.Password)
		if registerAccount(body.FirstName, body.LastName, body.Email, body.Password) {
			fmt.Println("Successfully Registered the account")
			context.JSON(http.StatusOK, true)
		} else {
			context.JSON(http.StatusBadRequest, false)
		}

	})
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
	h := sha256.New()
	h.Write([]byte(password))

	_, err = db.Exec(AUTH_INSERT, fname, lname, email, base64.URLEncoding.EncodeToString(h.Sum(nil)))
	if err != nil {
		panic(err) // Error Handling and sending back to the user is yet to be implemented.
		return false
	}

	fmt.Println("Successfully created account!")
	return true
}
func validateLogin(body LoginInfo) (ContactCard, bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=password dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	h := sha256.New()
	h.Write([]byte(body.Password))
	rows, err := db.Query(VALID_ACC, body.Email, base64.URLEncoding.EncodeToString(h.Sum(nil)))
	if err != nil {
		panic(err)
		return ContactCard{}, false
	}
	card := ContactCard{}

	for rows.Next() {
		err = rows.Scan(&card.FirstName, &card.LastName, &card.Email)
		if err != nil {
			panic(err)
		}

	}

	fmt.Println("The card is ", card)
	fmt.Println("Vetted account credentials")
	return card, true
}

func userExists(email string) bool {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=password dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
		return false
	}
	defer db.Close()
	_, err = db.Query(ACC_EXISTS, email)
	if err != nil {
		panic(err)
		return false
	}
	fmt.Println("Account exists")
	return true

}

func addContactDynamo(contact ContactCard) bool {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "us-east-2"
		return nil
	})

	svc := dynamodb.NewFromConfig(cfg)

	out, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(CONTACTS_TABLE),
		Item: map[string]types.AttributeValue{
			"email":     &types.AttributeValueMemberS{Value: contact.Email},
			"firstname": &types.AttributeValueMemberS{Value: contact.FirstName},
			"lastname":  &types.AttributeValueMemberS{Value: contact.LastName},
		},
	})
	fmt.Println(out)
	if err != nil {
		panic(err)
		return false
	}

	return true
}
