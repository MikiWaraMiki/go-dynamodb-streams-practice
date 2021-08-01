package infra_handler

import (
	"github.com/aws/aws-sdk-go/aws"
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
	session := session.Must(
		session.NewSessionWithOptions(session.Options{
			Profile: "localstack",
		}),
	)
	db := dynamo.New(session, &aws.Config{
		Region:     aws.String("us-east-2"),
		Endpoint:   aws.String("localhost:8000"),
		DisableSSL: aws.Bool(true),
	})

	return db
}

func generateProdDynamodbSession() *dynamo.DB {
	return dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
}
