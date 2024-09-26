package orders

import "api-gateway/entity/orders"

type Service interface {
	CreateOrder(order *orders.Order) error
	UserOrders(userID int32) ([]orders.Order, error)
	CallbackNotification(orderIDExt, status string, paidAmount int32) error
}
