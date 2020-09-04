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

type ZellerElements struct {
	Year	int	`json:"year"`
	Month	int	`json:"month"`
	Day		int	`json:"day"`
}