package transaction

type CreateTransactionPayload struct {
	CustomerID      string
	TransactionType string
	Amount          int64  `json:"amount" form:"amount"`
	ReferenceID     string `json:"reference_id" form:"reference_id"`
}
