package email

type Email struct {
	To              string `json:"to"`
	Subject         string `json:"subject"`
	Body            string `json:"body"`
	Type            string `json:"type"` // "verification" or "receipt"
	VerificationURL string `json:"verification_url,omitempty"`
}
