// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"context"
	"sync"
	"tracking-order-service/internal/app/domain"
)

// Ensure, that AnterajaRepositoryMock does implement AnterajaRepository.
// If this is not the case, regenerate this file with moq.
var _ AnterajaRepository = &AnterajaRepositoryMock{}

// AnterajaRepositoryMock is a mock implementation of AnterajaRepository.
//
// 	func TestSomethingThatUsesAnterajaRepository(t *testing.T) {
//
// 		// make and configure a mocked AnterajaRepository
// 		mockedAnterajaRepository := &AnterajaRepositoryMock{
// 			GetLastStatusOrderFunc: func(orderStatus string, anterajaApi domain.AnterajaResponseBodyApi) domain.ResultAnterajaStatus {
// 				panic("mock out the GetLastStatusOrder method")
// 			},
// 			TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultAnterajaRepository {
// 				panic("mock out the Tracking method")
// 			},
// 		}
//
// 		// use mockedAnterajaRepository in code that requires AnterajaRepository
// 		// and then make assertions.
//
// 	}
type AnterajaRepositoryMock struct {
	// GetLastStatusOrderFunc mocks the GetLastStatusOrder method.
	GetLastStatusOrderFunc func(orderStatus string, anterajaApi domain.AnterajaResponseBodyApi) domain.ResultAnterajaStatusOrder

	// TrackingFunc mocks the Tracking method.
	TrackingFunc func(ctx context.Context, config domain.RepositoryParam) domain.ResultAnterajaRepository

	// calls tracks calls to the methods.
	calls struct {
		// GetLastStatusOrder holds details about calls to the GetLastStatusOrder method.
		GetLastStatusOrder []struct {
			// OrderStatus is the orderStatus argument value.
			OrderStatus string
			// AnterajaApi is the anterajaApi argument value.
			AnterajaApi domain.AnterajaResponseBodyApi
		}
		// Tracking holds details about calls to the Tracking method.
		Tracking []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Config is the config argument value.
			Config domain.RepositoryParam
		}
	}
	lockGetLastStatusOrder sync.RWMutex
	lockTracking           sync.RWMutex
}

// GetLastStatusOrder calls GetLastStatusOrderFunc.
func (mock *AnterajaRepositoryMock) GetLastStatusOrder(orderStatus string, anterajaApi domain.AnterajaResponseBodyApi) domain.ResultAnterajaStatusOrder {
	if mock.GetLastStatusOrderFunc == nil {
		panic("AnterajaRepositoryMock.GetLastStatusOrderFunc: method is nil but AnterajaRepository.GetLastStatusOrder was just called")
	}
	callInfo := struct {
		OrderStatus string
		AnterajaApi domain.AnterajaResponseBodyApi
	}{
		OrderStatus: orderStatus,
		AnterajaApi: anterajaApi,
	}
	mock.lockGetLastStatusOrder.Lock()
	mock.calls.GetLastStatusOrder = append(mock.calls.GetLastStatusOrder, callInfo)
	mock.lockGetLastStatusOrder.Unlock()
	return mock.GetLastStatusOrderFunc(orderStatus, anterajaApi)
}

// GetLastStatusOrderCalls gets all the calls that were made to GetLastStatusOrder.
// Check the length with:
//     len(mockedAnterajaRepository.GetLastStatusOrderCalls())
func (mock *AnterajaRepositoryMock) GetLastStatusOrderCalls() []struct {
	OrderStatus string
	AnterajaApi domain.AnterajaResponseBodyApi
} {
	var calls []struct {
		OrderStatus string
		AnterajaApi domain.AnterajaResponseBodyApi
	}
	mock.lockGetLastStatusOrder.RLock()
	calls = mock.calls.GetLastStatusOrder
	mock.lockGetLastStatusOrder.RUnlock()
	return calls
}

// Tracking calls TrackingFunc.
func (mock *AnterajaRepositoryMock) Tracking(ctx context.Context, config domain.RepositoryParam) domain.ResultAnterajaRepository {
	if mock.TrackingFunc == nil {
		panic("AnterajaRepositoryMock.TrackingFunc: method is nil but AnterajaRepository.Tracking was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Config domain.RepositoryParam
	}{
		Ctx:    ctx,
		Config: config,
	}
	mock.lockTracking.Lock()
	mock.calls.Tracking = append(mock.calls.Tracking, callInfo)
	mock.lockTracking.Unlock()
	return mock.TrackingFunc(ctx, config)
}

// TrackingCalls gets all the calls that were made to Tracking.
// Check the length with:
//     len(mockedAnterajaRepository.TrackingCalls())
func (mock *AnterajaRepositoryMock) TrackingCalls() []struct {
	Ctx    context.Context
	Config domain.RepositoryParam
} {
	var calls []struct {
		Ctx    context.Context
		Config domain.RepositoryParam
	}
	mock.lockTracking.RLock()
	calls = mock.calls.Tracking
	mock.lockTracking.RUnlock()
	return calls
}