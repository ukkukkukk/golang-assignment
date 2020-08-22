package Struct

type Event struct {
	EventID    string `json:"id"`
	CustomerID string `json:"customer_id"`
	LoadAmount string `json:"load_amount"`
	EventTime  string `json:"time"`
}
