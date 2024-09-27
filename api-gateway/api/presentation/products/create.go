package products

type CreateRequest struct {
	Name  string `json:"name"`
	Price int32  `json:"price"`
}

type CreateDetailRequest struct {
	Size  int32 `json:"size"`
	Stock int32 `json:"stock"`
}
