package infra_handler

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func NewDynamoDBSession(environment string) *dynamo.DB {
	var db *dynamo.DB
	switch {
	case environment == "local":
		db = generateLocalDynamodbSession()
	case environment == "prod":
		db = generateProdDynamodbSession()
	}

	return db
}

func generateLocalDynamodbSession() *dynamo.DB {
	session, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-2"),
		Endpoint: aws.String("http://localhost:8000"),
		Credentials: credentials.NewStaticCredentials(
			"dummy",
			"dummy",
			"dummy",
		),
	})

	if err != nil {
		fmt.Printf("%v \n", err)
		panic(err)
	}

	db := dynamo.New(session)

	return db
}

func generateProdDynamodbSession() *dynamo.DB {
	return dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
}
