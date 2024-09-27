package products

import (
	PRODUCTS_ENTITY "api-gateway/entity/products"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ProductService struct {
	baseURL string
	client  *http.Client
}

func NewProductService() *ProductService {
	return &ProductService{
		baseURL: os.Getenv("PRODUCT_SERVICE_ADDR"),
		client:  http.DefaultClient,
	}
}

func (ps *ProductService) CreateProduct(product *PRODUCTS_ENTITY.Shoe) error {
	payload, err := json.Marshal(product)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", ps.baseURL+"/admin/shoe-models", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	res, err := ps.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(product)
	if err != nil {
		return err
	}
	log.Printf("%+v", product)
	return nil
}

func (ps *ProductService) CreateProductDetail(productDetail *PRODUCTS_ENTITY.ShoeDetail) error {
	payload, err := json.Marshal(productDetail)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/admin/shoe-models/%d", ps.baseURL, productDetail.ID),
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return err
	}
	res, err := ps.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(productDetail)
	if err != nil {
		return err
	}
	log.Printf("%+v", productDetail)
	return nil
}

func (ps *ProductService) UpdateProduct(product *PRODUCTS_ENTITY.Shoe) error {
	payload, err := json.Marshal(product)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/shoe-models/%d", ps.baseURL, product.ID), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	res, err := ps.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(product)
	if err != nil {
		return err
	}
	log.Printf("%+v", product)
	return nil
}

func (ps *ProductService) DeleteProduct(productID int32) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/admin/shoe-models/%d", ps.baseURL, productID), nil)
	if err != nil {
		return err
	}
	_, err = ps.client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProductService) GetAllShoes() ([]PRODUCTS_ENTITY.Shoe, error) {
	var products []PRODUCTS_ENTITY.Shoe
	req, err := http.NewRequest("GET", ps.baseURL+"/customer/products", nil)
	if err != nil {
		return nil, err
	}
	res, err := ps.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetShoeModelByID(modelID int32) (*PRODUCTS_ENTITY.Shoe, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/customer/products/%d", ps.baseURL, modelID), nil)
	if err != nil {
		return nil, err
	}
	res, err := ps.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	product := new(PRODUCTS_ENTITY.Shoe)
	err = json.NewDecoder(res.Body).Decode(&product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
