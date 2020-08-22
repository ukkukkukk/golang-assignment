package Functions

import (
	"encoding/json"
	"golang-assignment/Struct"
	"log"
)

func GenerateOutputRecord(event Struct.Event, accepted bool) string {
	var rejectedOutput = Struct.OutputEvent{CustomerID: event.CustomerID, EventID: event.EventID, Accepted: accepted}
	rejectedOutputJSON, jsonParseError := json.Marshal(rejectedOutput)

	if jsonParseError != nil {
		log.Fatal(jsonParseError)
		return ""
	}

	return string(rejectedOutputJSON)
}
