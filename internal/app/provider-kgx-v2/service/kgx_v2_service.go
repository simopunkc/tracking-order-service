package service

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/pkg/constant"
)

//go:generate moq -out kgx_v2_service_mock_test.go . KgxV2Repository
type KgxV2Repository interface {
	Tracking(context.Context, domain.RepositoryParam) domain.ResultKgxV2Repository
	GetLastStatusOrder(string, domain.KgxV2ResponseBodyApi) domain.ResultKgxV2StatusOrder
}

type KgxV2Service struct {
	kgxV2Repo KgxV2Repository
}

func NewKgxV2Service(kgxV2Repo KgxV2Repository) *KgxV2Service {
	return &KgxV2Service{kgxV2Repo}
}

var ctx = context.Background()

func (kvs KgxV2Service) RunTracking(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
	config := domain.RepositoryParam{
		Order:   order,
		Setting: rules,
		Direct:  true,
		Raw:     true,
	}

	respProvider := kvs.kgxV2Repo.Tracking(ctx, config)
	if respProvider.Error != nil {
		log.Println(order.Order.Awb, respProvider.Error)
		return domain.ResultProviderService{
			Hash: "GMnMvql7oSKY",
		}
	}

	lastCheckStatus := kvs.kgxV2Repo.GetLastStatusOrder(order.Order.Status, respProvider.ObjectResponse)
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
			Hash:     "GMoJpTgvbtk9",
			NewOrder: order,
		}
	}

	var orderTracking domain.KgxV2ResponseBodyApi
	err := json.Unmarshal([]byte(order.Response), &orderTracking)
	if err != nil {
		log.Println(order.Order.Awb, err)
		return domain.ResultProviderService{
			Hash: "GM5lVgaxG0yE",
		}
	}

	if len(respProvider.ObjectResponse.Data) != len(orderTracking.Data) {
		order.Response = respProvider.RawResponse
		order.ModifiedAt = time.Now()
		return domain.ResultProviderService{
			Hash:     "GMhJNDAjEEBa",
			NewOrder: order,
		}
	}

	if order.ModifiedAt.AddDate(0, 0, rules.MaxDayTracking).Unix() < time.Now().Unix() {
		order.Order.Status = constant.NOT_TRACKABLE
		order.Order.ModifiedAt = time.Now()
		return domain.ResultProviderService{
			Hash:     "GM8e1M64zsQu",
			NewOrder: order,
		}
	}

	return domain.ResultProviderService{
		Hash: "GMrkgsMZmRGp",
	}
}
