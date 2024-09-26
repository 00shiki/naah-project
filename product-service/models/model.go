package models

import "time"

type Voucher struct {
	Voucher_id string  `gorm:"primaryKey"`
	Discount   float64 `gorm:"type:decimal(10,2);not null"`
	ExpiryDate time.Time
	Used       bool `gorm:"not null"`
}

type User struct {
	UserID       uint      `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Email        string    `gorm:"unique;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	FirstName    string    `gorm:"not null" json:"first_name"`
	LastName     string    `gorm:"not null" json:"last_name"`
	BirthDate    time.Time `gorm:"type:date" json:"birth_date,omitempty"`
	Address      string    `gorm:"type:text" json:"address,omitempty"`
	ContactNo    string    `gorm:"size:20" json:"contact_no,omitempty"`
}

type ShoeModel struct {
	ModelID int    `gorm:"primaryKey;autoIncrement"`
	Name    string `gorm:"not null"`
	Price   int    `gorm:"not null"`
}

type ShoeDetail struct {
	ShoeID  int `gorm:"primaryKey;autoIncrement"`
	ModelID int `gorm:"not null"`
	Size    int `gorm:"not null"`
	Stock   int `gorm:"not null"`
}

type Cart struct {
	CartID   int `gorm:"primaryKey;autoIncrement"`
	UserID   int `gorm:"not null"`
	Quantity int `gorm:"not null"`
	ShoeID   int `gorm:"not null"`
}

type Order struct {
	OrderID    int `gorm:"primaryKey;autoIncrement"`
	UserID     int `gorm:"not null"`
	VoucherID  *uint
	Status     string  `gorm:"not null"`
	Price      int     `gorm:"not null"`
	Fee        int     `gorm:"not null"`
	Discount   float64 `gorm:"type:decimal(10,2)"`
	TotalPrice int     `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Metadata   string `gorm:"type:text"`
}

type Payment struct {
	PaymentID         int    `gorm:"primaryKey;autoIncrement"`
	OrderID           int    `gorm:"not null"`
	PaymentExternalID string `gorm:"size:36"`
	Amount            int    `gorm:"not null"`
	Status            string `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Metadata          string `gorm:"type:text"`
}

type Delivery struct {
	DeliveryID        int       `gorm:"primaryKey;autoIncrement"`
	OrderID           int       `gorm:"not null"`
	DeliveryDate      time.Time `gorm:"type:date;not null"`
	ArrivalDate       *time.Time
	CourierName       string
	CourierService    string
	WeightGrams       int `gorm:"not null"`
	OriginCityID      string
	DestinationCityID string
	DeliveryFee       int    `gorm:"not null"`
	Status            string `gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Metadata          string `gorm:"type:text"`
}

type OrderDetail struct {
	OrderDetailID int `gorm:"primaryKey;autoIncrement"`
	OrderID       int `gorm:"not null"`
	ShoeID        int `gorm:"not null"`
	Quantity      int `gorm:"not null"`
}
