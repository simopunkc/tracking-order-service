package service

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/pkg/constant"
)

//go:generate moq -out anteraja_service_mock_test.go . AnterajaRepository
type AnterajaRepository interface {
	Tracking(context.Context, domain.RepositoryParam) domain.ResultAnterajaRepository
	GetLastStatusOrder(string, domain.AnterajaResponseBodyApi) domain.ResultAnterajaStatusOrder
}

type AnterajaService struct {
	anterajaRepo AnterajaRepository
}

func NewAnterajaService(anterajaRepo AnterajaRepository) *AnterajaService {
	return &AnterajaService{anterajaRepo}
}

var ctx = context.Background()

func (as AnterajaService) RunTracking(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
	config := domain.RepositoryParam{
		Order:   order,
		Setting: rules,
		Direct:  true,
		Raw:     true,
	}

	respProvider := as.anterajaRepo.Tracking(ctx, config)
	if respProvider.Error != nil {
		log.Println(order.Order.Awb, respProvider.Error)
		return domain.ResultProviderService{
			Hash: "GMBEJkbD6SPm",
		}
	}

	if respProvider.ObjectResponse.Status != 200 {
		log.Println(order.Order.Awb, "status not 200", respProvider.RawResponse)
		return domain.ResultProviderService{
			Hash: "GMT4bPm9ecBA",
		}
	}

	lastCheckStatus := as.anterajaRepo.GetLastStatusOrder(order.Order.Status, respProvider.ObjectResponse)
	if lastCheckStatus.ModifiedAt.After(order.ModifiedAt) {
		order.Order.Status = lastCheckStatus.Status
		order.Order.ModifiedAt = time.Now()
		order.ModifiedAt = lastCheckStatus.ModifiedAt
	}

	if order.OrderID == 0 {
		order.OrderID = order.Order.ID
		order.Response = respProvider.RawResponse
		order.ModifiedAt = time.Now()
		return domain.ResultProviderService{
			Hash:     "GMv19TqabRlJ",
			NewOrder: order,
		}
	}

	var orderTracking domain.AnterajaResponseBodyApi
	err := json.Unmarshal([]byte(order.Response), &orderTracking)
	if err != nil {
		log.Println(order.Order.Awb, err)
		return domain.ResultProviderService{
			Hash: "GMWpVLpeTJgG",
		}
	}

	if len(respProvider.ObjectResponse.Content.History) != len(orderTracking.Content.History) {
		order.Response = respProvider.RawResponse
		order.ModifiedAt = time.Now()
		return domain.ResultProviderService{
			Hash:     "GMXWyuXNnMGg",
			NewOrder: order,
		}
	}

	if order.ModifiedAt.AddDate(0, 0, rules.MaxDayTracking).Unix() < time.Now().Unix() {
		order.Order.Status = constant.NOT_TRACKABLE
		order.Order.ModifiedAt = time.Now()
		return domain.ResultProviderService{
			Hash:     "GM8Ebv74zWl2",
			NewOrder: order,
		}
	}

	return domain.ResultProviderService{
		Hash: "GMSjB3s2Dvot",
	}
}
