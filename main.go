package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	/*var awsConfig = aws.Config{
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
	var dynamoDBClient = dynamodb.New(awsSession)*/

	var fileName = "input.txt"

	log.Printf("Reading from input file: %s \n", fileName)

	var inputFile, fileOpenError = os.Open(fileName)
	if fileOpenError != nil {
		log.Fatal(fileOpenError)
		return
	}

	var fileScanner = bufio.NewScanner(inputFile)

	var linesProcessed = 0

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		linesProcessed++;
	}

	log.Printf("Processed %d lines.\n", linesProcessed)

	var scannerError = fileScanner.Err()

	if scannerError != nil {
		log.Fatal(scannerError)
	}

	var fileCloseError = inputFile.Close()
	if fileCloseError != nil {
		log.Fatal(fileCloseError)
	}
}
