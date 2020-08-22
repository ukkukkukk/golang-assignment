package Functions

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"golang-assignment/Struct"
	"log"
	"time"
)

func ValidateEvent(dynamoDBClient *dynamodb.DynamoDB, event Struct.Event) {
	var correlationRecord = QueryEventCorrelationTable(dynamoDBClient, "EventCorrelationTable", event.CustomerID, event.EventID)

	if correlationRecord == nil {
		log.Fatal("Correlation query failed. ")
		return
	}

	if *correlationRecord.Count > 0 {
		//we have seen this record before, output false
		var rejectedOutput = Struct.OutputEvent{CustomerID: event.CustomerID, EventID: event.EventID, Accepted: "false"}
		rejectedOutputJSON, jsonParseError := json.Marshal(rejectedOutput)

		if jsonParseError != nil {
			log.Fatal(jsonParseError)
			return
		}
		log.Println(string(rejectedOutputJSON))
	}

	//extract date from dateTime
	var layout = "2006-01-02T15:04:05Z"
	var eventTime, timestampParseError = time.Parse(layout, event.EventTime)

	if timestampParseError != nil {
		log.Fatal(timestampParseError)
		return
	}

	var extractedEventDate = eventTime.Format("2006-01-02")

	//find date for start of week

	var extractedStartOfWeek = FindStartOfWeek(eventTime)

	var loadTableRecords = QueryCustomerLoadTable(dynamoDBClient, "DailyCustomerLoadTable", event.CustomerID, extractedStartOfWeek, extractedEventDate)

	if loadTableRecords == nil {
		log.Fatal("Load  table query failed. ")
		return
	}

	fmt.Println(correlationRecord)
	fmt.Println(loadTableRecords)

}
