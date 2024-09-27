package products

type UpdateRequest struct {
	Name  string `json:"name"`
	Price int32  `json:"price"`
}
