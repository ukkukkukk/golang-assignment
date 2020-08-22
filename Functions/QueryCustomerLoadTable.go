package Functions

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

func QueryCustomerLoadTable(dynamoDBClient *dynamodb.DynamoDB, tableName string, hashKeyValue string, rangeKeyValue1 string, rangeKeyValue2 string) {

	var queryParam = &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("#CUSTOMER_ID = :CUSTOMER_ID AND #DATE BETWEEN :START_DATE AND :END_DATE"),
		ExpressionAttributeNames: map[string]*string{
			"#CUSTOMER_ID": aws.String("CUSTOMER_ID"),
			"#DATE":        aws.String("DATE"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":CUSTOMER_ID": {
				S: aws.String(hashKeyValue),
			},
			":START_DATE": {
				S: aws.String(rangeKeyValue1),
			},
			":END_DATE": {
				S: aws.String(rangeKeyValue2),
			},
		},
	}

	resp, queryError := dynamoDBClient.Query(queryParam)
	if queryError != nil {
		log.Fatal(queryError)
		return
	}

	fmt.Println(queryParam)
	fmt.Println(resp)

}
