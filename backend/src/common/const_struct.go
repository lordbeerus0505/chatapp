package common

const (
	HOST           = "localhost"
	PORT           = 5432
	USER           = "abhiram"
	DBNAME         = "chatapp"
	DYNAMOREGION   = "us-east-2"
	CONTACTS_TABLE = "contacts"
)

type ContactCard struct {
	Email     string
	FirstName string
	LastName  string
	User      string
}
