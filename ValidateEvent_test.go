package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"golang-assignment/Functions"
	"golang-assignment/Struct"
	"log"
	"testing"
)

//must have dynamodb running locally when running these tests
//need to delete and recreate tables before running this test
//ideally the dynamodb calls should be mocked, assertions should be used

func TestValidateEvent_AcceptedCaseSimple(t *testing.T) {
	var awsConfig = aws.Config{
		Region:   aws.String("us-east-2"),
		Endpoint: aws.String("http://localhost:8000"),
	}
	log.Printf("Creating AWS session. \n")
	var awsSession, sessionError = session.NewSession(&awsConfig)

	if sessionError != nil {
		log.Fatal(sessionError)
		return
	}

	log.Printf("Creating DynamoDB Client. \n")
	var dynamoDBClient = dynamodb.New(awsSession)

	var event = Struct.Event{"1a", "5a", "$5.00", "2018-01-01T01:02:00Z"}

	var output = Functions.ValidateEvent(dynamoDBClient, event)
	var expectedOutput = "{\"id\":\"1a\",\"customer_id\":\"5a\",\"accepted\":true}"

	if output != expectedOutput {
		t.Errorf("Output does not match expected output.")
	}
}

func TestValidateEvent_DuplicateEvent(t *testing.T) {
	var awsConfig = aws.Config{
		Region:   aws.String("us-east-2"),
		Endpoint: aws.String("http://localhost:8000"),
	}
	log.Printf("Creating AWS session. \n")
	var awsSession, sessionError = session.NewSession(&awsConfig)

	if sessionError != nil {
		log.Fatal(sessionError)
		return
	}

	log.Printf("Creating DynamoDB Client. \n")
	var dynamoDBClient = dynamodb.New(awsSession)

	var event = Struct.Event{"2a", "6a", "$5.00", "2018-01-01T01:02:00Z"}

	var output = Functions.ValidateEvent(dynamoDBClient, event)
	var expectedOutput = "{\"id\":\"2a\",\"customer_id\":\"6a\",\"accepted\":true}"

	if output != expectedOutput {
		t.Errorf("Output does not match expected output.")
	}

	output = Functions.ValidateEvent(dynamoDBClient, event)

	if output != "" {
		t.Errorf("Output does not match expected output.")
	}

}

func TestValidateEvent_ExceedDailyLimit(t *testing.T) {
	var awsConfig = aws.Config{
		Region:   aws.String("us-east-2"),
		Endpoint: aws.String("http://localhost:8000"),
	}
	log.Printf("Creating AWS session. \n")
	var awsSession, sessionError = session.NewSession(&awsConfig)

	if sessionError != nil {
		log.Fatal(sessionError)
		return
	}

	log.Printf("Creating DynamoDB Client. \n")
	var dynamoDBClient = dynamodb.New(awsSession)

	var event = Struct.Event{"3a", "7a", "$3000.99", "2018-01-01T01:02:00Z"}

	var output = Functions.ValidateEvent(dynamoDBClient, event)
	var expectedOutput = "{\"id\":\"3a\",\"customer_id\":\"7a\",\"accepted\":true}"

	if output != expectedOutput {
		t.Errorf("Output does not match expected output.")
	}

	var event2 = Struct.Event{"4a", "7a", "$2000.01", "2018-01-01T02:02:00Z"}
	output = Functions.ValidateEvent(dynamoDBClient, event2)
	var expectedOutput2 = "{\"id\":\"4a\",\"customer_id\":\"7a\",\"accepted\":false}"

	if output != expectedOutput2 {
		t.Errorf("Output does not match expected output.")
	}

	var event3 = Struct.Event{"5a", "7a", "$1999.01", "2018-01-01T03:02:00Z"}
	output = Functions.ValidateEvent(dynamoDBClient, event3)
	var expectedOutput3 = "{\"id\":\"5a\",\"customer_id\":\"7a\",\"accepted\":true}"

	if output != expectedOutput3 {
		t.Errorf("Output does not match expected output.")
	}

}

func TestValidateEvent_ExceedDailyLoadsLimit(t *testing.T) {
	var awsConfig = aws.Config{
		Region:   aws.String("us-east-2"),
		Endpoint: aws.String("http://localhost:8000"),
	}
	log.Printf("Creating AWS session. \n")
	var awsSession, sessionError = session.NewSession(&awsConfig)

	if sessionError != nil {
		log.Fatal(sessionError)
		return
	}

	log.Printf("Creating DynamoDB Client. \n")
	var dynamoDBClient = dynamodb.New(awsSession)

	var event = Struct.Event{"3a", "8a", "$1.00", "2018-01-01T01:02:00Z"}

	var output = Functions.ValidateEvent(dynamoDBClient, event)
	var expectedOutput = "{\"id\":\"3a\",\"customer_id\":\"8a\",\"accepted\":true}"

	if output != expectedOutput {
		t.Errorf("Output does not match expected output.")
	}

	var event2 = Struct.Event{"4a", "8a", "$1.00", "2018-01-01T02:02:00Z"}
	output = Functions.ValidateEvent(dynamoDBClient, event2)
	var expectedOutput2 = "{\"id\":\"4a\",\"customer_id\":\"8a\",\"accepted\":true}"

	if output != expectedOutput2 {
		t.Errorf("Output does not match expected output.")
	}

	var event3 = Struct.Event{"5a", "8a", "$1.00", "2018-01-01T03:02:00Z"}
	output = Functions.ValidateEvent(dynamoDBClient, event3)
	var expectedOutput3 = "{\"id\":\"5a\",\"customer_id\":\"8a\",\"accepted\":true}"

	if output != expectedOutput3 {
		t.Errorf("Output does not match expected output.")
	}

	var event4 = Struct.Event{"6a", "8a", "$1.00", "2018-01-01T04:02:00Z"}
	output = Functions.ValidateEvent(dynamoDBClient, event4)
	var expectedOutput4 = "{\"id\":\"6a\",\"customer_id\":\"8a\",\"accepted\":false}"

	if output != expectedOutput4 {
		t.Errorf("Output does not match expected output.")
	}

}

func TestValidateEvent_ExceedWeeklyLimit(t *testing.T) {
	var awsConfig = aws.Config{
		Region:   aws.String("us-east-2"),
		Endpoint: aws.String("http://localhost:8000"),
	}
	log.Printf("Creating AWS session. \n")
	var awsSession, sessionError = session.NewSession(&awsConfig)

	if sessionError != nil {
		log.Fatal(sessionError)
		return
	}

	log.Printf("Creating DynamoDB Client. \n")
	var dynamoDBClient = dynamodb.New(awsSession)

	var event = Struct.Event{"3a", "9a", "$5000.00", "2020-08-17T01:02:00Z"}

	var output = Functions.ValidateEvent(dynamoDBClient, event)
	var expectedOutput = "{\"id\":\"3a\",\"customer_id\":\"9a\",\"accepted\":true}"

	if output != expectedOutput {
		t.Errorf("Output does not match expected output.")
	}

	var event2 = Struct.Event{"4a", "9a", "$5000.00", "2020-08-19T02:02:00Z"}
	output = Functions.ValidateEvent(dynamoDBClient, event2)
	var expectedOutput2 = "{\"id\":\"4a\",\"customer_id\":\"9a\",\"accepted\":true}"

	if output != expectedOutput2 {
		t.Errorf("Output does not match expected output.")
	}

	var event3 = Struct.Event{"5a", "9a", "$5000.00", "2020-08-20T03:02:00Z"}
	output = Functions.ValidateEvent(dynamoDBClient, event3)
	var expectedOutput3 = "{\"id\":\"5a\",\"customer_id\":\"9a\",\"accepted\":true}"

	if output != expectedOutput3 {
		t.Errorf("Output does not match expected output.")
	}

	var event4 = Struct.Event{"6a", "9a", "$5000.00", "2020-08-22T04:02:00Z"}
	output = Functions.ValidateEvent(dynamoDBClient, event4)
	var expectedOutput4 = "{\"id\":\"6a\",\"customer_id\":\"9a\",\"accepted\":true}"

	if output != expectedOutput4 {
		t.Errorf("Output does not match expected output.")
	}

	var event5 = Struct.Event{"7a", "9a", "$5000.00", "2020-08-23T05:02:00Z"}
	output = Functions.ValidateEvent(dynamoDBClient, event5)
	var expectedOutput5 = "{\"id\":\"7a\",\"customer_id\":\"9a\",\"accepted\":false}"

	if output != expectedOutput5 {
		t.Errorf("Output does not match expected output.")
	}

	var event6 = Struct.Event{"8a", "9a", "$5000.00", "2020-08-24T05:02:00Z"}
	output = Functions.ValidateEvent(dynamoDBClient, event6)
	var expectedOutput6 = "{\"id\":\"8a\",\"customer_id\":\"9a\",\"accepted\":true}"

	if output != expectedOutput6 {
		t.Errorf("Output does not match expected output.")
	}

}
