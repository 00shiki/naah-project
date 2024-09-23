package model

type EmailPayload struct {
    To      		string `json:"to"`
    Subject 		string `json:"subject"`
    Type    		string `json:"type"` // "verification" or "receipt"
	VerificationURL string `json:"verification_url,omitempty"`
    OrderID       	string `json:"order_id,omitempty"`
}