package service

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/pkg/constant"
)

//go:generate moq -out sicepat_service_mock_test.go . SicepatRepository
type SicepatRepository interface {
	Tracking(context.Context, domain.RepositoryParam) domain.ResultSicepatRepository
	GetLastStatusOrder(string, domain.SicepatResponseBodyApi) domain.ResultSicepatStatusOrder
}

type SicepatService struct {
	sicepatRepo SicepatRepository
}

func NewSicepatService(sicepatRepo SicepatRepository) *SicepatService {
	return &SicepatService{sicepatRepo}
}

var ctx = context.Background()

func (ss SicepatService) RunTracking(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
	config := domain.RepositoryParam{
		Order:   order,
		Setting: rules,
		Direct:  true,
		Raw:     true,
	}

	respProvider := ss.sicepatRepo.Tracking(ctx, config)
	if respProvider.Error != nil {
		log.Println(order.Order.Awb, respProvider.Error)
		return domain.ResultProviderService{
			Hash: "GMmNattIDtUg",
		}
	}

	if respProvider.ObjectResponse.Sicepat.Status.Code != 200 {
		log.Println(order.Order.Awb, "status not 200", respProvider.RawResponse)
		return domain.ResultProviderService{
			Hash: "GM9z3BeFFaUT",
		}
	}

	lastCheckStatus := ss.sicepatRepo.GetLastStatusOrder(order.Order.Status, respProvider.ObjectResponse)
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
			Hash:     "GMvzuHX2PKva",
			NewOrder: order,
		}
	}

	var orderTracking domain.SicepatResponseBodyApi
	err := json.Unmarshal([]byte(order.Response), &orderTracking)
	if err != nil {
		log.Println(order.Order.Awb, err)
		return domain.ResultProviderService{
			Hash: "GMw6XJsF5TiU",
		}
	}

	if len(respProvider.ObjectResponse.Sicepat.Result.TrackHistory) != len(orderTracking.Sicepat.Result.TrackHistory) {
		order.Response = respProvider.RawResponse
		order.ModifiedAt = time.Now()
		return domain.ResultProviderService{
			Hash:     "GMsSBuMBnzuG",
			NewOrder: order,
		}
	}

	if respProvider.ObjectResponse.Sicepat.Result.LastStatus.DateTime != orderTracking.Sicepat.Result.LastStatus.DateTime {
		order.Response = respProvider.RawResponse
		order.ModifiedAt = time.Now()
		return domain.ResultProviderService{
			Hash:     "GMvM0YohP1lu",
			NewOrder: order,
		}
	}

	if order.ModifiedAt.AddDate(0, 0, rules.MaxDayTracking).Unix() < time.Now().Unix() {
		order.Order.Status = constant.NOT_TRACKABLE
		order.Order.ModifiedAt = time.Now()
		return domain.ResultProviderService{
			Hash:     "GMGQ1HTnw2vs",
			NewOrder: order,
		}
	}

	return domain.ResultProviderService{
		Hash: "GM5j2RzZRraO",
	}
}
