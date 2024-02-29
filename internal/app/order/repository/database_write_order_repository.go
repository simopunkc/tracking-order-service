package repository

import (
	"tracking-order-service/internal/app/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DatabaseWriteOrderRepository struct {
	db *gorm.DB
}

func NewDatabaseWriteOrderRepository(db *gorm.DB) *DatabaseWriteOrderRepository {
	return &DatabaseWriteOrderRepository{db}
}

func (or DatabaseWriteOrderRepository) BulkUpdateOrder(orders []domain.Order) (int64, error) {
	bulkInsertOrUpdate := or.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"status", "modified_at"}),
	}).Omit(clause.Associations).Create(&orders)

	return bulkInsertOrUpdate.RowsAffected, bulkInsertOrUpdate.Error
}

func (or DatabaseWriteOrderRepository) BulkUpdateOrderTracking(orders []domain.OrderTracking) (int64, error) {
	bulkInsertOrUpdate := or.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "order_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"response", "modified_at"}),
	}).Omit(clause.Associations).Create(&orders)

	return bulkInsertOrUpdate.RowsAffected, bulkInsertOrUpdate.Error
}
