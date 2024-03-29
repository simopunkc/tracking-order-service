// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"sync"
	"tracking-order-service/internal/app/domain"
)

// Ensure, that DatabaseReadOrderRepositoryMock does implement DatabaseReadOrderRepository.
// If this is not the case, regenerate this file with moq.
var _ DatabaseReadOrderRepository = &DatabaseReadOrderRepositoryMock{}

// DatabaseReadOrderRepositoryMock is a mock implementation of DatabaseReadOrderRepository.
//
// 	func TestSomethingThatUsesDatabaseReadOrderRepository(t *testing.T) {
//
// 		// make and configure a mocked DatabaseReadOrderRepository
// 		mockedDatabaseReadOrderRepository := &DatabaseReadOrderRepositoryMock{
// 			OrdersModelFunc: func(trackingOrderParam domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
// 				panic("mock out the OrdersModel method")
// 			},
// 			RulesModelFunc: func(trackingOrderParam domain.TrackingOrderParam) (domain.Settings, error) {
// 				panic("mock out the RulesModel method")
// 			},
// 		}
//
// 		// use mockedDatabaseReadOrderRepository in code that requires DatabaseReadOrderRepository
// 		// and then make assertions.
//
// 	}
type DatabaseReadOrderRepositoryMock struct {
	// OrdersModelFunc mocks the OrdersModel method.
	OrdersModelFunc func(trackingOrderParam domain.TrackingOrderParam) ([]domain.OrderTracking, error)

	// RulesModelFunc mocks the RulesModel method.
	RulesModelFunc func(trackingOrderParam domain.TrackingOrderParam) (domain.Settings, error)

	// calls tracks calls to the methods.
	calls struct {
		// OrdersModel holds details about calls to the OrdersModel method.
		OrdersModel []struct {
			// TrackingOrderParam is the trackingOrderParam argument value.
			TrackingOrderParam domain.TrackingOrderParam
		}
		// RulesModel holds details about calls to the RulesModel method.
		RulesModel []struct {
			// TrackingOrderParam is the trackingOrderParam argument value.
			TrackingOrderParam domain.TrackingOrderParam
		}
	}
	lockOrdersModel sync.RWMutex
	lockRulesModel  sync.RWMutex
}

// OrdersModel calls OrdersModelFunc.
func (mock *DatabaseReadOrderRepositoryMock) OrdersModel(trackingOrderParam domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
	if mock.OrdersModelFunc == nil {
		panic("DatabaseReadOrderRepositoryMock.OrdersModelFunc: method is nil but DatabaseReadOrderRepository.OrdersModel was just called")
	}
	callInfo := struct {
		TrackingOrderParam domain.TrackingOrderParam
	}{
		TrackingOrderParam: trackingOrderParam,
	}
	mock.lockOrdersModel.Lock()
	mock.calls.OrdersModel = append(mock.calls.OrdersModel, callInfo)
	mock.lockOrdersModel.Unlock()
	return mock.OrdersModelFunc(trackingOrderParam)
}

// OrdersModelCalls gets all the calls that were made to OrdersModel.
// Check the length with:
//     len(mockedDatabaseReadOrderRepository.OrdersModelCalls())
func (mock *DatabaseReadOrderRepositoryMock) OrdersModelCalls() []struct {
	TrackingOrderParam domain.TrackingOrderParam
} {
	var calls []struct {
		TrackingOrderParam domain.TrackingOrderParam
	}
	mock.lockOrdersModel.RLock()
	calls = mock.calls.OrdersModel
	mock.lockOrdersModel.RUnlock()
	return calls
}

// RulesModel calls RulesModelFunc.
func (mock *DatabaseReadOrderRepositoryMock) RulesModel(trackingOrderParam domain.TrackingOrderParam) (domain.Settings, error) {
	if mock.RulesModelFunc == nil {
		panic("DatabaseReadOrderRepositoryMock.RulesModelFunc: method is nil but DatabaseReadOrderRepository.RulesModel was just called")
	}
	callInfo := struct {
		TrackingOrderParam domain.TrackingOrderParam
	}{
		TrackingOrderParam: trackingOrderParam,
	}
	mock.lockRulesModel.Lock()
	mock.calls.RulesModel = append(mock.calls.RulesModel, callInfo)
	mock.lockRulesModel.Unlock()
	return mock.RulesModelFunc(trackingOrderParam)
}

// RulesModelCalls gets all the calls that were made to RulesModel.
// Check the length with:
//     len(mockedDatabaseReadOrderRepository.RulesModelCalls())
func (mock *DatabaseReadOrderRepositoryMock) RulesModelCalls() []struct {
	TrackingOrderParam domain.TrackingOrderParam
} {
	var calls []struct {
		TrackingOrderParam domain.TrackingOrderParam
	}
	mock.lockRulesModel.RLock()
	calls = mock.calls.RulesModel
	mock.lockRulesModel.RUnlock()
	return calls
}

// Ensure, that DatabaseWriteOrderRepositoryMock does implement DatabaseWriteOrderRepository.
// If this is not the case, regenerate this file with moq.
var _ DatabaseWriteOrderRepository = &DatabaseWriteOrderRepositoryMock{}

// DatabaseWriteOrderRepositoryMock is a mock implementation of DatabaseWriteOrderRepository.
//
// 	func TestSomethingThatUsesDatabaseWriteOrderRepository(t *testing.T) {
//
// 		// make and configure a mocked DatabaseWriteOrderRepository
// 		mockedDatabaseWriteOrderRepository := &DatabaseWriteOrderRepositoryMock{
// 			BulkUpdateOrderFunc: func(orders []domain.Order) (int64, error) {
// 				panic("mock out the BulkUpdateOrder method")
// 			},
// 			BulkUpdateOrderTrackingFunc: func(orders []domain.OrderTracking) (int64, error) {
// 				panic("mock out the BulkUpdateOrderTracking method")
// 			},
// 		}
//
// 		// use mockedDatabaseWriteOrderRepository in code that requires DatabaseWriteOrderRepository
// 		// and then make assertions.
//
// 	}
type DatabaseWriteOrderRepositoryMock struct {
	// BulkUpdateOrderFunc mocks the BulkUpdateOrder method.
	BulkUpdateOrderFunc func(orders []domain.Order) (int64, error)

	// BulkUpdateOrderTrackingFunc mocks the BulkUpdateOrderTracking method.
	BulkUpdateOrderTrackingFunc func(orders []domain.OrderTracking) (int64, error)

	// calls tracks calls to the methods.
	calls struct {
		// BulkUpdateOrder holds details about calls to the BulkUpdateOrder method.
		BulkUpdateOrder []struct {
			// Orders is the orders argument value.
			Orders []domain.Order
		}
		// BulkUpdateOrderTracking holds details about calls to the BulkUpdateOrderTracking method.
		BulkUpdateOrderTracking []struct {
			// Orders is the orders argument value.
			Orders []domain.OrderTracking
		}
	}
	lockBulkUpdateOrder         sync.RWMutex
	lockBulkUpdateOrderTracking sync.RWMutex
}

// BulkUpdateOrder calls BulkUpdateOrderFunc.
func (mock *DatabaseWriteOrderRepositoryMock) BulkUpdateOrder(orders []domain.Order) (int64, error) {
	if mock.BulkUpdateOrderFunc == nil {
		panic("DatabaseWriteOrderRepositoryMock.BulkUpdateOrderFunc: method is nil but DatabaseWriteOrderRepository.BulkUpdateOrder was just called")
	}
	callInfo := struct {
		Orders []domain.Order
	}{
		Orders: orders,
	}
	mock.lockBulkUpdateOrder.Lock()
	mock.calls.BulkUpdateOrder = append(mock.calls.BulkUpdateOrder, callInfo)
	mock.lockBulkUpdateOrder.Unlock()
	return mock.BulkUpdateOrderFunc(orders)
}

// BulkUpdateOrderCalls gets all the calls that were made to BulkUpdateOrder.
// Check the length with:
//     len(mockedDatabaseWriteOrderRepository.BulkUpdateOrderCalls())
func (mock *DatabaseWriteOrderRepositoryMock) BulkUpdateOrderCalls() []struct {
	Orders []domain.Order
} {
	var calls []struct {
		Orders []domain.Order
	}
	mock.lockBulkUpdateOrder.RLock()
	calls = mock.calls.BulkUpdateOrder
	mock.lockBulkUpdateOrder.RUnlock()
	return calls
}

// BulkUpdateOrderTracking calls BulkUpdateOrderTrackingFunc.
func (mock *DatabaseWriteOrderRepositoryMock) BulkUpdateOrderTracking(orders []domain.OrderTracking) (int64, error) {
	if mock.BulkUpdateOrderTrackingFunc == nil {
		panic("DatabaseWriteOrderRepositoryMock.BulkUpdateOrderTrackingFunc: method is nil but DatabaseWriteOrderRepository.BulkUpdateOrderTracking was just called")
	}
	callInfo := struct {
		Orders []domain.OrderTracking
	}{
		Orders: orders,
	}
	mock.lockBulkUpdateOrderTracking.Lock()
	mock.calls.BulkUpdateOrderTracking = append(mock.calls.BulkUpdateOrderTracking, callInfo)
	mock.lockBulkUpdateOrderTracking.Unlock()
	return mock.BulkUpdateOrderTrackingFunc(orders)
}

// BulkUpdateOrderTrackingCalls gets all the calls that were made to BulkUpdateOrderTracking.
// Check the length with:
//     len(mockedDatabaseWriteOrderRepository.BulkUpdateOrderTrackingCalls())
func (mock *DatabaseWriteOrderRepositoryMock) BulkUpdateOrderTrackingCalls() []struct {
	Orders []domain.OrderTracking
} {
	var calls []struct {
		Orders []domain.OrderTracking
	}
	mock.lockBulkUpdateOrderTracking.RLock()
	calls = mock.calls.BulkUpdateOrderTracking
	mock.lockBulkUpdateOrderTracking.RUnlock()
	return calls
}

// Ensure, that ProviderServiceMock does implement ProviderService.
// If this is not the case, regenerate this file with moq.
var _ ProviderService = &ProviderServiceMock{}

// ProviderServiceMock is a mock implementation of ProviderService.
//
// 	func TestSomethingThatUsesProviderService(t *testing.T) {
//
// 		// make and configure a mocked ProviderService
// 		mockedProviderService := &ProviderServiceMock{
// 			RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
// 				panic("mock out the RunTracking method")
// 			},
// 		}
//
// 		// use mockedProviderService in code that requires ProviderService
// 		// and then make assertions.
//
// 	}
type ProviderServiceMock struct {
	// RunTrackingFunc mocks the RunTracking method.
	RunTrackingFunc func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService

	// calls tracks calls to the methods.
	calls struct {
		// RunTracking holds details about calls to the RunTracking method.
		RunTracking []struct {
			// Order is the order argument value.
			Order domain.OrderTracking
			// Rules is the rules argument value.
			Rules *domain.SettingsValueCheckOrder
		}
	}
	lockRunTracking sync.RWMutex
}

// RunTracking calls RunTrackingFunc.
func (mock *ProviderServiceMock) RunTracking(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
	if mock.RunTrackingFunc == nil {
		panic("ProviderServiceMock.RunTrackingFunc: method is nil but ProviderService.RunTracking was just called")
	}
	callInfo := struct {
		Order domain.OrderTracking
		Rules *domain.SettingsValueCheckOrder
	}{
		Order: order,
		Rules: rules,
	}
	mock.lockRunTracking.Lock()
	mock.calls.RunTracking = append(mock.calls.RunTracking, callInfo)
	mock.lockRunTracking.Unlock()
	return mock.RunTrackingFunc(order, rules)
}

// RunTrackingCalls gets all the calls that were made to RunTracking.
// Check the length with:
//     len(mockedProviderService.RunTrackingCalls())
func (mock *ProviderServiceMock) RunTrackingCalls() []struct {
	Order domain.OrderTracking
	Rules *domain.SettingsValueCheckOrder
} {
	var calls []struct {
		Order domain.OrderTracking
		Rules *domain.SettingsValueCheckOrder
	}
	mock.lockRunTracking.RLock()
	calls = mock.calls.RunTracking
	mock.lockRunTracking.RUnlock()
	return calls
}
