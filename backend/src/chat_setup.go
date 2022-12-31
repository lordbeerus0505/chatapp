package src

import (
	"backend/src/common"
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
)

const ACC_EXISTS = "SELECT COUNT(*) FROM auth WHERE email=$1"

type ContactsList struct {
	Email    string               `dynamodbav:"email"`
	Contacts []common.ContactCard `dynamodbav:"contacts"`
}

func AddContact(r *gin.Engine) {
	r.POST("/add-contact", func(context *gin.Context) {
		body := common.ContactCard{}
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

func GetChats(r *gin.Engine) {
	// Could be made a GET request as well
	r.POST("/get-chats", func(context *gin.Context) {
		body := common.ContactCard{} // only need email but receiving all; future proofing
		if err := context.BindJSON(&body); err != nil {
			context.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(body)
		// Note, Context.TODO() can only be used in local methods not exported ones (no caps)
		status, listOfChats := retrieveContacts(body.Email)
		if !status {
			context.JSON(http.StatusBadRequest, ContactsList{Email: "False"})
		} else {
			context.JSON(http.StatusOK, &listOfChats)
		}

	})
}

func addContactDynamo(contact common.ContactCard) bool {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = common.DYNAMOREGION
		return nil
	})

	svc := dynamodb.NewFromConfig(cfg)
	fmt.Println("SVC is", svc)
	// Retrieve contacts from Dynamo, and then append to the list.
	var listOfChats ContactsList
	status, listOfChats := retrieveContacts(contact.User)
	if !status {
		listOfChats = ContactsList{
			Email: contact.User,
			Contacts: []common.ContactCard{
				contact,
			},
		}
	} else {
		listOfChats.Contacts = append(listOfChats.Contacts, contact)
	}

	av, err := attributevalue.MarshalMap(listOfChats)
	fmt.Println(av, listOfChats)
	if err != nil {
		panic(err)
	}
	out, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(common.CONTACTS_TABLE),
		Item:      av,
	})
	fmt.Println(out)
	if err != nil {
		panic(err)
		return false
	}

	return true
}

func userExists(email string) bool {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=password dbname=%s sslmode=disable",
		common.HOST, common.PORT, common.USER, common.DBNAME)
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

func retrieveContacts(email string) (bool, ContactsList) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = common.DYNAMOREGION
		return nil
	})
	svc := dynamodb.NewFromConfig(cfg)
	filt := expression.Name("email").Equal(expression.Value(email))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	out, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(common.CONTACTS_TABLE),
	})

	if err != nil {
		panic(err)
	}

	var contacts []ContactsList
	err = attributevalue.UnmarshalListOfMaps(out.Items, &contacts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Items in Dynamo are", contacts)
	if len(contacts) != 1 {
		return false, ContactsList{}
	}
	return true, contacts[0]
}
