package Struct

type OutputEvent struct {
	EventID    string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}
