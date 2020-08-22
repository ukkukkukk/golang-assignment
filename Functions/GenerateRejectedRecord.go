package Functions

import (
	"encoding/json"
	"golang-assignment/Struct"
	"log"
)

func GenerateRejectedRecord(event Struct.Event) string {
	var rejectedOutput = Struct.OutputEvent{CustomerID: event.CustomerID, EventID: event.EventID, Accepted: "false"}
	rejectedOutputJSON, jsonParseError := json.Marshal(rejectedOutput)

	if jsonParseError != nil {
		log.Fatal(jsonParseError)
		return ""
	}

	return string(rejectedOutputJSON)
}
