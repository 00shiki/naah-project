package models

import "time"

type Voucher struct {
	VoucherID  string    `json:"voucher_id"`
	Discount   float64   `json:"discount"`
	ExpiryDate time.Time `json:"expiry_date"`
	Used       bool      `json:"used"`
}

type User struct {
	UserID       uint      `json:"user_id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	BirthDate    time.Time `json:"birth_date,omitempty"`
	Address      string    `json:"address,omitempty"`
	ContactNo    string    `json:"contact_no,omitempty"`
}

type ShoeModel struct {
	ModelID int    `json:"model_id"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
}

type ShoeDetail struct {
	ShoeID  int `json:"shoe_id"`
	ModelID int `json:"model_id"`
	Size    int `json:"size"`
	Stock   int `json:"stock"`
}

type UpdateShoeDetailRequest struct {
	ShoeDetail ShoeDetail `json:"shoe_detail"`
	ShoeModel  ShoeModel  `json:"shoe_model"`
}

type Cart struct {
	CartID   int `json:"cart_id"`
	UserID   int `json:"user_id"`
	Quantity int `json:"quantity"`
	ShoeID   int `json:"shoe_id"`
}

type Order struct {
	OrderID    int       `json:"order_id"`
	UserID     int       `json:"user_id"`
	VoucherID  *string   `json:"voucher_id,omitempty"` // Use pointer for optional value
	Status     string    `json:"status"`
	Price      int       `json:"price"`
	Fee        int       `json:"fee"`
	Discount   float64   `json:"discount,omitempty"`
	TotalPrice int       `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Metadata   string    `json:"metadata,omitempty"`
}

type Payment struct {
	PaymentID         int       `json:"payment_id"`
	OrderID           int       `json:"order_id"`
	PaymentExternalID string    `json:"payment_external_id,omitempty"`
	Amount            int       `json:"amount"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Metadata          string    `json:"metadata,omitempty"`
}

type Delivery struct {
	DeliveryID        int        `json:"delivery_id"`
	OrderID           int        `json:"order_id"`
	DeliveryDate      time.Time  `json:"delivery_date"`
	ArrivalDate       *time.Time `json:"arrival_date,omitempty"`
	CourierName       string     `json:"courier_name,omitempty"`
	CourierService    string     `json:"courier_service,omitempty"`
	WeightGrams       int        `json:"weight_grams"`
	OriginCityID      string     `json:"origin_city_id,omitempty"`
	DestinationCityID string     `json:"destination_city_id,omitempty"`
	DeliveryFee       int        `json:"delivery_fee"`
	Status            string     `json:"status"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	Metadata          string     `json:"metadata,omitempty"`
}

type OrderDetail struct {
	OrderDetailID int `json:"order_detail_id"`
	OrderID       int `json:"order_id"`
	ShoeID        int `json:"shoe_id"`
	Quantity      int `json:"quantity"`
}
