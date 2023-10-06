package models

type Transaction struct {
	TransactionId int32   `json:"transactionId"`
	UserId        int32   `json:"userId"`
	Date          string  `json:"date"`
	Transaction   float64 `json:"transaction"`
}
