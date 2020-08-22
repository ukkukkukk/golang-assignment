package Functions

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"golang-assignment/Struct"
	"log"
)

func InsertLoadTableRecord(dynamoDBClient *dynamodb.DynamoDB, tableName string, record Struct.LoadTableRecord) {

	var marshalledRecord, marshalError = dynamodbattribute.MarshalMap(record)

	if marshalError != nil {
		log.Fatal(marshalError)
		return
	}

	var putParam = &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      marshalledRecord,
	}

	var _, err = dynamoDBClient.PutItem(putParam)

	if err != nil {
		log.Fatal(err)
		return
	}

	return

}
