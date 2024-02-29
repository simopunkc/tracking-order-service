package service

import (
	"context"
	"log"
	"tracking-order-service/internal/app/domain"
)

//go:generate moq -out ninja_service_mock_test.go . NinjaRepository
type NinjaRepository interface {
	Tracking(context.Context, domain.RepositoryParam) domain.ResultNinjaRepository
	GetLastStatusOrder([]domain.NinjaResponseBodyApi, domain.NinjaResponseBodyApi) domain.ListNinjaResponseBodyApi
}

type NinjaService struct {
	ninjaRepo NinjaRepository
}

func NewNinjaService(ninjaRepo NinjaRepository) *NinjaService {
	return &NinjaService{ninjaRepo}
}

var ctx = context.Background()

func (ns NinjaService) RunTracking(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
	config := domain.RepositoryParam{
		Order:   order,
		Setting: rules,
		Direct:  true,
		Raw:     true,
	}

	respProvider := ns.ninjaRepo.Tracking(ctx, config)
	if respProvider.Error != nil {
		log.Println(order.Order.Awb, respProvider.Error)
		return domain.ResultProviderService{
			Hash: "GMxsmc9q1het",
		}
	}

	return domain.ResultProviderService{
		Hash: "GMdiUEk2E7R8",
	}
}
