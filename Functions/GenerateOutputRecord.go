package Functions

import (
	"encoding/json"
	"golang-assignment/Struct"
	"log"
)

func GenerateOutputRecord(event Struct.Event, accepted bool) string {
	var outputRecord = Struct.OutputEvent{CustomerID: event.CustomerID, EventID: event.EventID, Accepted: accepted}
	outputJSON, jsonParseError := json.Marshal(outputRecord)

	if jsonParseError != nil {
		log.Fatal(jsonParseError)
		return ""
	}

	return string(outputJSON)
}
