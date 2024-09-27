package products

type Shoe struct {
	ID          int32        `json:"model_id"`
	Name        string       `json:"name"`
	Price       int32        `json:"price"`
	ShoeDetails []ShoeDetail `json:"shoe_details,omitempty"`
}

type ShoeDetail struct {
	ID      int32 `json:"shoe_id"`
	Stock   int32 `json:"stock"`
	Size    int32 `json:"size"`
	ModelID int32 `json:"model_id,omitempty"`
	Shoe    `json:"-"`
}
