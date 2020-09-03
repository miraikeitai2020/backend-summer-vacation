package model

type User struct {
	Name	string `json:"name"`
}

// TASK1 RESPONSE BODY
type TimeStamp struct {
	TimeStamp	string	`json:"timestamp"`
	Detail		Detail	`json:"datail"`
}

type Detail struct {
	Date	string	`json:"date"`
	Time	string	`json:"time"`
}