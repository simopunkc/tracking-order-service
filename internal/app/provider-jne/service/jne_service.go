package service

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/pkg/constant"
)

//go:generate moq -out jne_service_mock_test.go . JneRepository
type JneRepository interface {
	Tracking(context.Context, domain.RepositoryParam) domain.ResultJneRepository
	GetLastStatusOrder(string, domain.JneResponseBodyApi) domain.ResultJneStatusOrder
}

type JneService struct {
	jneRepo JneRepository
}

func NewJneService(jneRepo JneRepository) *JneService {
	return &JneService{jneRepo}
}

var ctx = context.Background()

func (js JneService) RunTracking(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
	config := domain.RepositoryParam{
		Order:   order,
		Setting: rules,
		Direct:  true,
		Raw:     true,
	}

	respProvider := js.jneRepo.Tracking(ctx, config)
	if respProvider.Error != nil {
		log.Println(order.Order.Awb, respProvider.Error)
		return domain.ResultProviderService{
			Hash: "GMs6XQBsuOM2",
		}
	}

	lastCheckStatus := js.jneRepo.GetLastStatusOrder(order.Order.Status, respProvider.ObjectResponse)
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
			Hash:     "GMzOwFgFFknl",
			NewOrder: order,
		}
	}

	var orderTracking domain.JneResponseBodyApi
	err := json.Unmarshal([]byte(order.Response), &orderTracking)
	if err != nil {
		log.Println(order.Order.Awb, err)
		return domain.ResultProviderService{
			Hash: "GMlGlTZjc6QX",
		}
	}

	if len(respProvider.ObjectResponse.History) != len(orderTracking.History) {
		order.Response = respProvider.RawResponse
		order.ModifiedAt = time.Now()
		return domain.ResultProviderService{
			Hash:     "GMUDom1i573D",
			NewOrder: order,
		}
	}

	if order.ModifiedAt.AddDate(0, 0, rules.MaxDayTracking).Unix() < time.Now().Unix() {
		order.Order.Status = constant.NOT_TRACKABLE
		order.Order.ModifiedAt = time.Now()
		return domain.ResultProviderService{
			Hash:     "GMJZHTLf97RC",
			NewOrder: order,
		}
	}

	return domain.ResultProviderService{
		Hash: "GMwuJeC7TRHn",
	}
}
