package orders

import (
	"api-gateway/entity/orders"
	"api-gateway/entity/products"
	pb "api-gateway/proto"
	"context"
	"time"
)

type OrderService struct {
	client pb.OrderServiceClient
}

func NewOrderService(client pb.OrderServiceClient) *OrderService {
	return &OrderService{
		client: client,
	}
}

func (os *OrderService) CreateOrder(order *orders.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cartIds := make([]int32, len(order.Delivery.Carts))
	for i, cart := range order.Delivery.Carts {
		cartIds[i] = cart.CartID
	}
	req := &pb.AddOrderRequest{
		UserId:             order.UserID,
		VoucherId:          order.VoucherID,
		CartIds:            cartIds,
		CourierName:        order.Delivery.Courier.Name,
		CourierServiceName: order.Delivery.Courier.Service,
		OriginCityId:       order.Delivery.OriginCityID,
		DestinationCityId:  order.Delivery.DestinationCityID,
		OtherFee:           order.Fee,
		Metadata:           order.Metadata,
	}
	res, err := os.client.AddOrder(ctx, req)
	if err != nil {
		return err
	}
	order.TotalPrice = res.TotalPrice
	order.InvoiceUrl = res.InvoiceUrl
	order.ExpiredDate = res.ExpiredDate
	return nil
}

func (os *OrderService) UserOrders(userID int32) ([]orders.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.GetOrderListRequest{
		UserId: userID,
	}
	res, err := os.client.GetOrderList(ctx, req)
	if err != nil {
		return nil, err
	}
	userOrders := make([]orders.Order, len(res.Orders))
	for i, order := range res.Orders {
		orderItems := make([]products.Product, len(order.Shoes))
		for j, shoe := range order.Shoes {
			orderItems[j] = products.Product{
				Name:  shoe.Name,
				Price: shoe.Price,
				Stock: shoe.Qty,
				Size:  shoe.Size,
			}
		}
		userOrders[i] = orders.Order{
			ID:         order.OrderId,
			OrderItems: orderItems,
			Fee:        order.Fee,
			Discount:   order.Discount,
			TotalPrice: order.TotalPrice,
			VoucherID:  order.VoucherId,
			Status:     order.Status,
		}
	}
	return userOrders, nil
}

func (os *OrderService) CallbackNotification(orderIDExt, status string, paidAmount int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.CallbackNotificationRequest{
		OrderIdExt: orderIDExt,
		Status:     status,
		PaidAmount: paidAmount,
	}
	_, err := os.client.CallbackNotification(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
