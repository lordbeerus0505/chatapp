# Chat Application
### Frontend
#### Requirements - 
Install React, NodeJS, npm
npm install react-router-dom
npm install axios
npm install @hookform/resolvers
npm install yup
npm install react-hook-form
### Backend
#### Requirements - 
Install Go.
To setup gin use https://stackoverflow.com/questions/72972764/golang-no-required-module-provides-package
go get github.com/rs/cors/wrapper/gin - for cors
go get -u github.com/lib/pq - postgresql
go get "github.com/aws/aws-sdk-go-v2"
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/dynamodb
go get github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression
go get github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue
### Database
#### PostgreSQL
Install PostgreSQL
Run Auth.SQL
#### DynamoDB
##### Contacts
Create table with email as PK in DynamoDB. Create a user in AWS. Get the access keys from it.
Add the access keys to environment variables https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/
