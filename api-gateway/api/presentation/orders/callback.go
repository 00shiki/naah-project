package orders

type CallbackRequest struct {
	PaidAmount int32  `json:"paid_amount"`
	Status     string `json:"status"`
	ExternalID string `json:"external_id"`
}
