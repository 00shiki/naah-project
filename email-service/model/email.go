package model

type OrderDetail struct {
    ShoeName     string `json:"shoe_name"`
    ShoeSize     int64  `json:"shoe_size"`
    Qty          int32  `json:"quantity"`
}

type OrderReceiptEmail struct {
    UserName   string `json:"user_name"`
    UserEmail   string `json:"user_email"`
    OrderId     int64  `json:"order_id"`
    TotalPrice  int32  `json:"total_price"`
    OrderDate   string `json:"order_date"`
    OrderDetail []OrderDetail `json:"order_detail"`
}

type EmailPayload struct {
    To      		string `json:"to"`
    Subject 		string `json:"subject"`
    Type    		string `json:"type"` // "verification" or "receipt" or "delivered" or "verified"
	VerificationURL string `json:"verification_url,omitempty"`
    OrderReceipt    OrderReceiptEmail `json:"order_receipt,omitempty"`
    OrderID         int64 `json:"order_id,omitempty"`
}