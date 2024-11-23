package wallet

type WalletResponse struct {
	ID         string `json:"id"`
	OwnedBy    string `json:"owned_by"`
	Status     string `json:"status"`
	EnabledAt  string `json:"enabled_at,omitempty"`
	DisabledAt string `json:"disabled_at,omitempty"`
	Balance    int64  `json:"balance"`
}
