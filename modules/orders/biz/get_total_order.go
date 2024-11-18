package biz

import (
	"context"
	"first-app/modules/orders/models"
	"time"
)

type GetTotalOrderStorage interface {
	GetTotalOrder(ctx context.Context) ([]models.OrderGet,error)
}

type GetTotalOrderBiz struct {
	store GetTotalOrderStorage
}

func NewGetTotalOrderByIdBiz(store GetTotalOrderStorage) *GetTotalOrderBiz {
	return &GetTotalOrderBiz{store: store}
}

func (biz *GetTotalOrderBiz) GetNewTotalOrder(ctx context.Context) (*models.TotalOrder, error) {
	var data []models.OrderGet
	totalOrder := &models.TotalOrder{}

	data, err := biz.store.GetTotalOrder(ctx)
	if err != nil {
		return nil, err
	}

	today := time.Now().Format("2006-01-02")
	currentMonth := time.Now().Format("2006-01")

	for _, order := range data {

		orderDate := order.CreatedAt.Format("2006-01-02")
		orderMonth := order.CreatedAt.Format("2006-01")

		orderRevenue := float64(order.Quantity) * order.Menu.Price

		if orderDate == today {
			totalOrder.TotalRevenueByDay += orderRevenue
			totalOrder.TotelOrderByDay += order.Quantity
		}

		if orderMonth == currentMonth {
			totalOrder.TotalRevenueByMonth += orderRevenue
			totalOrder.TotelOrderByMonth += order.Quantity
		}
	}

	return totalOrder, nil
}
