package customer

type InitCustomerPayload struct {
	CustomerXID string `json:"customer_xid" form:"customer_xid" binding:"required"`
}
