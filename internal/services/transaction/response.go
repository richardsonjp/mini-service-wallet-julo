package transaction

type DepositResponse struct {
	ID          string `json:"id"`
	DepositedBy string `json:"deposited_by"`
	Status      string `json:"status"`
	DepositedAt string `json:"deposited_at,omitempty"`
	Amount      int64  `json:"amount"`
	ReferenceID string `json:"reference_id"`
}

type WithdrawalResponse struct {
	ID          string `json:"id"`
	WithdrawBy  string `json:"withdraw_by_by"`
	Status      string `json:"status"`
	WithdrawAt  string `json:"withdraw_at_at,omitempty"`
	Amount      int64  `json:"amount"`
	ReferenceID string `json:"reference_id"`
}

type ListTransactionsResponse struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	TransactedAt string `json:"transacted_at"`
	Type         string `json:"type"`
	Amount       int64  `json:"amount"`
	ReferenceID  string `json:"reference_id"`
}
