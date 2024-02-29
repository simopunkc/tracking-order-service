package service

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/pkg/constant"
)

//go:generate moq -out pos_service_mock_test.go . PosRepository
type PosRepository interface {
	Tracking(context.Context, domain.RepositoryParam) domain.ResultPosRepository
	GetLastStatusOrder(string, domain.PosResponseBodyApi) domain.ResultPosStatusOrder
}

type PosService struct {
	posRepo PosRepository
}

func NewPosService(posRepo PosRepository) *PosService {
	return &PosService{posRepo}
}

var ctx = context.Background()

func (ps PosService) RunTracking(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
	config := domain.RepositoryParam{
		Order:   order,
		Setting: rules,
		Direct:  true,
		Raw:     true,
	}

	respProvider := ps.posRepo.Tracking(ctx, config)
	if respProvider.Error != nil {
		log.Println(order.Order.Awb, respProvider.Error)
		return domain.ResultProviderService{
			Hash: "GMVdeWwgd80S",
		}
	}

	lastCheckStatus := ps.posRepo.GetLastStatusOrder(order.Order.Status, respProvider.ObjectResponse)
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
			Hash:     "GMlIEjkOokSJ",
			NewOrder: order,
		}
	}

	var orderTracking domain.PosResponseBodyApi
	err := json.Unmarshal([]byte(order.Response), &orderTracking)
	if err != nil {
		log.Println(order.Order.Awb, err)
		return domain.ResultProviderService{
			Hash: "GMyO1jQpZfh8",
		}
	}

	if orderTracking.RsTnt.RTnt != nil && orderTracking.Response.Data == nil {
		respProvider.ObjectResponse.Response.Data = respProvider.ObjectResponse.RsTnt.RTnt
	}

	if len(respProvider.ObjectResponse.Response.Data) != len(orderTracking.Response.Data) {
		order.Response = respProvider.RawResponse
		order.ModifiedAt = time.Now()
		return domain.ResultProviderService{
			Hash:     "GMyvln2PQbxV",
			NewOrder: order,
		}
	}

	if order.ModifiedAt.AddDate(0, 0, rules.MaxDayTracking).Unix() < time.Now().Unix() {
		order.Order.Status = constant.NOT_TRACKABLE
		order.Order.ModifiedAt = time.Now()
		return domain.ResultProviderService{
			Hash:     "GMjdDjebsCox",
			NewOrder: order,
		}
	}

	return domain.ResultProviderService{
		Hash: "GM8dDIHq3aUw",
	}
}
