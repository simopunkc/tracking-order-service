package handler

import (
	"fmt"
	"tracking-order-service/internal/app/domain"
)

type OrderService interface {
	RunCronJob(domain.TrackingOrderParam) domain.ResultOrderService
}

type OrderHandler struct {
	orderService OrderService
}

func NewOrderHandler(orderService OrderService) *OrderHandler {
	return &OrderHandler{orderService}
}

func (oh OrderHandler) PrintCronJob(argument domain.TrackingOrderParam) {
	result := oh.orderService.RunCronJob(argument)
	fmt.Println(result.TotalUpdated, "models updated")
}
