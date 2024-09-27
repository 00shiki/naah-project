package products

import PRODUCTS_ENTITY "api-gateway/entity/products"

type Service interface {
	CreateProduct(product *PRODUCTS_ENTITY.Shoe) error
	CreateProductDetail(productDetail *PRODUCTS_ENTITY.ShoeDetail) error
	UpdateProduct(product *PRODUCTS_ENTITY.Shoe) error
	DeleteProduct(productID int32) error
	GetAllShoes() ([]PRODUCTS_ENTITY.Shoe, error)
	GetShoeModelByID(modelID int32) (*PRODUCTS_ENTITY.Shoe, error)
}
