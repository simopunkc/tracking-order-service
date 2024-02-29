package service

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/pkg/constant"
)

//go:generate moq -out jnt_service_mock_test.go . JntRepository
type JntRepository interface {
	Tracking(context.Context, domain.RepositoryParam) domain.ResultJntRepository
	GetLastStatusOrder(string, domain.JntResponseBodyApi) domain.ResultJntStatusOrder
}

type JntService struct {
	jntRepo JntRepository
}

func NewJntService(jntRepo JntRepository) *JntService {
	return &JntService{jntRepo}
}

var ctx = context.Background()

func (js JntService) RunTracking(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
	config := domain.RepositoryParam{
		Order:   order,
		Setting: rules,
		Direct:  true,
		Raw:     true,
	}

	respProvider := js.jntRepo.Tracking(ctx, config)
	if respProvider.Error != nil {
		log.Println(order.Order.Awb, respProvider.Error)
		return domain.ResultProviderService{
			Hash: "GMOa8O8UJskH",
		}
	}

	lastCheckStatus := js.jntRepo.GetLastStatusOrder(order.Order.Status, respProvider.ObjectResponse)
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
			Hash:     "GM2zqs0JxDAk",
			NewOrder: order,
		}
	}

	var orderTracking domain.JntResponseBodyApi
	err := json.Unmarshal([]byte(order.Response), &orderTracking)
	if err != nil {
		log.Println(order.Order.Awb, err)
		return domain.ResultProviderService{
			Hash: "GMZ2HpjpIK8K",
		}
	}

	if len(respProvider.ObjectResponse.History) != len(orderTracking.History) {
		order.Response = respProvider.RawResponse
		order.ModifiedAt = time.Now()
		return domain.ResultProviderService{
			Hash:     "GMLng7IsolKR",
			NewOrder: order,
		}
	}

	if order.ModifiedAt.AddDate(0, 0, rules.MaxDayTracking).Unix() < time.Now().Unix() {
		order.Order.Status = constant.NOT_TRACKABLE
		order.Order.ModifiedAt = time.Now()
		return domain.ResultProviderService{
			Hash:     "GMquingxcWCe",
			NewOrder: order,
		}
	}

	return domain.ResultProviderService{
		Hash: "GMtlsUlxC7SI",
	}
}
