package repository

import (
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/pkg/constant"

	"gorm.io/gorm"
)

type DatabaseReadOrderRepository struct {
	db *gorm.DB
}

func NewDatabaseReadOrderRepository(db *gorm.DB) *DatabaseReadOrderRepository {
	return &DatabaseReadOrderRepository{db}
}

func (or DatabaseReadOrderRepository) OrdersModel(argument domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
	var orders []domain.OrderTracking
	var filterQuery []string

	modelOrder := or.db.Select("*").Table("order_tracking ot").
		Joins("left join \"order\" o on ot.order_id = o.id").
		Joins("left join provider p on o.provider_id = p.id").
		Preload("Order").
		Preload("Order.Provider").
		Where("p.slug <> ?", constant.NINJA)

	if argument.OrderNumber != "" {
		modelOrder = modelOrder.Where("o.order_number = ?", argument.OrderNumber)
	}

	if !argument.BypassCheckStatus {
		filterQuery = append(filterQuery, constant.WAITING_PICKUP)
		filterQuery = append(filterQuery, constant.SENDING)
		filterQuery = append(filterQuery, constant.ON_COURIER)
	}
	if argument.NotTrackableOnly {
		filterQuery = append(filterQuery, constant.NOT_TRACKABLE)
	}

	if len(filterQuery) > 0 {
		filterQueryString := ""
		for i, status := range filterQuery {
			if i == 0 {
				filterQueryString += "'" + status + "'"
			} else {
				filterQueryString += ",'" + status + "'"
			}
		}
		modelOrder = modelOrder.Where("o.status in (" + filterQueryString + ")")
	}

	if argument.StartDate != "" {
		modelOrder = modelOrder.Where("o.created_at >= ?", argument.StartDate)
	}

	err := modelOrder.Where("o.awb <> ''").Find(&orders).Error

	return orders, err
}

func (or DatabaseReadOrderRepository) RulesModel(argument domain.TrackingOrderParam) (domain.Settings, error) {
	var rule domain.Settings
	err := or.db.Table("settings").
		Where("name = ?", "check_order").
		First(&rule).Error

	return rule, err
}
