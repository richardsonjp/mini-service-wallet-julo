package wallet

type ToggleWallet struct {
	CustomerID string `json:"customer_id" form:"customer_id"`
	Enabled    bool   `json:"enabled" form:"enabled"`
}
