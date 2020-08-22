package Functions

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func QueryEventCorrelationTable(dynamoDBClient *dynamodb.DynamoDB, tableName string, hashKeyValue string, rangeKeyValue string) {

	var queryParam = &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("#CUSTOMER_ID = :CUSTOMER_ID AND #EVENT_ID = :EVENT_ID"),
		ExpressionAttributeNames: map[string]*string{
			"#CUSTOMER_ID": aws.String("CUSTOMER_ID"),
			"#EVENT_ID":    aws.String("EVENT_ID"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":CUSTOMER_ID": {
				S: aws.String(hashKeyValue),
			},
			":EVENT_ID": {
				S: aws.String(rangeKeyValue),
			},
		},
	}

	resp, queryError := dynamoDBClient.Query(queryParam)
	if queryError != nil {
		log.Fatal(queryError)
		return
	}

	fmt.Println(resp)

}
