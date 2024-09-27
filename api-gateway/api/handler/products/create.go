package products

import (
	PRODUCTS_PRESENTATION "api-gateway/api/presentation/products"
	"api-gateway/entity/products"
	"api-gateway/entity/responses"
	"api-gateway/entity/users"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (handler *Controller) Create(c echo.Context) error {
	role, ok := c.Get("role").(float64)
	if !ok {
		res := &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: "Invalid Token",
		}
		return responses.HandleResponse(c, res)
	}

	if users.Role(role) != users.Admin {
		res := &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: "Restricted Access",
		}
		return responses.HandleResponse(c, res)
	}

	payload := new(PRODUCTS_PRESENTATION.CreateRequest)
	if err := c.Bind(payload); err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	product := &products.Shoe{
		Name:  payload.Name,
		Price: payload.Price,
	}
	err := handler.ps.CreateProduct(product)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}
	res := &responses.Response{
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    product,
	}
	return responses.HandleResponse(c, res)
}

func (handler *Controller) CreateDetail(c echo.Context) error {
	role, ok := c.Get("role").(float64)
	if !ok {
		res := &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: "Invalid Token",
		}
		return responses.HandleResponse(c, res)
	}

	if users.Role(role) != users.Admin {
		res := &responses.Response{
			Code:    http.StatusUnauthorized,
			Message: "Restricted Access",
		}
		return responses.HandleResponse(c, res)
	}

	productIDStr := c.Param("productID")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	payload := new(PRODUCTS_PRESENTATION.CreateDetailRequest)
	if err := c.Bind(payload); err != nil {
		res := &responses.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}

	productDetail := &products.ShoeDetail{
		Shoe: products.Shoe{
			ID: int32(productID),
		},
		Size:  payload.Size,
		Stock: payload.Stock,
	}
	err = handler.ps.CreateProductDetail(productDetail)
	if err != nil {
		res := &responses.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return responses.HandleResponse(c, res)
	}
	res := &responses.Response{
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    productDetail,
	}
	return responses.HandleResponse(c, res)
}
