package Functions

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"golang-assignment/Struct"
	"log"
	"strconv"
	"time"
)

const eventCorrelationTableName = "EventCorrelationTable"
const dailyCustomerLoadTableName = "DailyCustomerLoadTable"

func ValidateEvent(dynamoDBClient *dynamodb.DynamoDB, event Struct.Event) string {
	var correlationRecord = QueryEventCorrelationTable(dynamoDBClient, eventCorrelationTableName, event.CustomerID, event.EventID)

	if correlationRecord == nil {
		log.Fatal("Correlation query failed. ")
	}

	if *correlationRecord.Count > 0 {
		//we have seen this record before, ignore it
		return ""
	}

	var correlationRecordToInsert = Struct.CorrelationTableRecord{CUSTOMER_ID: event.CustomerID, EVENT_ID: event.EventID}
	InsertCorrelationTableRecord(dynamoDBClient, eventCorrelationTableName, correlationRecordToInsert)

	//extract date from dateTime
	var layout = "2006-01-02T15:04:05Z"
	var eventTime, timestampParseError = time.Parse(layout, event.EventTime)

	if timestampParseError != nil {
		log.Fatal(timestampParseError)
	}

	var extractedEventDate = eventTime.Format("2006-01-02")

	//find date for start of week

	var extractedStartOfWeek = FindStartOfWeek(eventTime)

	var loadTableRecords = QueryCustomerLoadTable(dynamoDBClient, dailyCustomerLoadTableName, event.CustomerID, extractedStartOfWeek, extractedEventDate)

	if loadTableRecords == nil {
		log.Fatal("Load  table query failed. ")
	}

	var currentAcceptedDailyLoadAmount = 0.0
	var currentAcceptedWeeklyLoadAmount = 0.0
	var currentAcceptedDailyLoads = 0

	//process accepted customer loads for the week
	for _, i := range loadTableRecords.Items {
		var item = Struct.LoadTableRecord{}

		var unmarshalError = dynamodbattribute.UnmarshalMap(i, &item)

		if unmarshalError != nil {
			log.Fatal(unmarshalError)
		}

		var dailyLoadAmount, parseError = strconv.ParseFloat(item.DAILY_LOAD_AMOUNT, 64)

		if parseError != nil {
			log.Fatal(parseError)
		}

		currentAcceptedWeeklyLoadAmount = currentAcceptedWeeklyLoadAmount + dailyLoadAmount

		if item.DATE == extractedEventDate {
			currentAcceptedDailyLoadAmount = dailyLoadAmount

			currentAcceptedDailyLoads, parseError = strconv.Atoi(item.NUMBER_OF_LOADS)
			if parseError != nil {
				log.Fatal(parseError)
			}
		}

	}

	var eventLoadAmount, parseError = strconv.ParseFloat(event.LoadAmount[1:], 64)

	if parseError != nil {
		log.Fatal(parseError)
	}

	var newDailyLoadAmount = currentAcceptedDailyLoadAmount + eventLoadAmount
	var newDailyLoads = currentAcceptedDailyLoads + 1
	var newWeeklyLoadAmount = currentAcceptedWeeklyLoadAmount + eventLoadAmount

	//check new daily/weekly amounts
	if newDailyLoadAmount > 5000.00 || newDailyLoads > 3 || newWeeklyLoadAmount > 20000.00 {
		var outputRecord = GenerateOutputRecord(event, false)
		log.Println(outputRecord)
		return outputRecord
	}

	//this is an accepted event

	var convertedNewDailyAmount = strconv.FormatFloat(newDailyLoadAmount, 'f', 2, 64)
	var convertedNewDailyLoads = strconv.Itoa(newDailyLoads)
	var loadRecordToInsert = Struct.LoadTableRecord{CUSTOMER_ID: event.CustomerID, DATE: extractedEventDate, DAILY_LOAD_AMOUNT: convertedNewDailyAmount, NUMBER_OF_LOADS: convertedNewDailyLoads}

	InsertLoadTableRecord(dynamoDBClient, dailyCustomerLoadTableName, loadRecordToInsert)

	var outputRecord = GenerateOutputRecord(event, true)
	log.Println(outputRecord)
	return outputRecord

}
