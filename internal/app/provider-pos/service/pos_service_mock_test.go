// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"context"
	"sync"
	"tracking-order-service/internal/app/domain"
)

// Ensure, that PosRepositoryMock does implement PosRepository.
// If this is not the case, regenerate this file with moq.
var _ PosRepository = &PosRepositoryMock{}

// PosRepositoryMock is a mock implementation of PosRepository.
//
// 	func TestSomethingThatUsesPosRepository(t *testing.T) {
//
// 		// make and configure a mocked PosRepository
// 		mockedPosRepository := &PosRepositoryMock{
// 			GetLastStatusOrderFunc: func(orderStatus string, posApi domain.PosResponseBodyApi) domain.ResultPosStatus {
// 				panic("mock out the GetLastStatusOrder method")
// 			},
// 			TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultPosRepository {
// 				panic("mock out the Tracking method")
// 			},
// 		}
//
// 		// use mockedPosRepository in code that requires PosRepository
// 		// and then make assertions.
//
// 	}
type PosRepositoryMock struct {
	// GetLastStatusOrderFunc mocks the GetLastStatusOrder method.
	GetLastStatusOrderFunc func(orderStatus string, posApi domain.PosResponseBodyApi) domain.ResultPosStatusOrder

	// TrackingFunc mocks the Tracking method.
	TrackingFunc func(ctx context.Context, config domain.RepositoryParam) domain.ResultPosRepository

	// calls tracks calls to the methods.
	calls struct {
		// GetLastStatusOrder holds details about calls to the GetLastStatusOrder method.
		GetLastStatusOrder []struct {
			// OrderStatus is the orderStatus argument value.
			OrderStatus string
			// PosApi is the posApi argument value.
			PosApi domain.PosResponseBodyApi
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
func (mock *PosRepositoryMock) GetLastStatusOrder(orderStatus string, posApi domain.PosResponseBodyApi) domain.ResultPosStatusOrder {
	if mock.GetLastStatusOrderFunc == nil {
		panic("PosRepositoryMock.GetLastStatusOrderFunc: method is nil but PosRepository.GetLastStatusOrder was just called")
	}
	callInfo := struct {
		OrderStatus string
		PosApi      domain.PosResponseBodyApi
	}{
		OrderStatus: orderStatus,
		PosApi:      posApi,
	}
	mock.lockGetLastStatusOrder.Lock()
	mock.calls.GetLastStatusOrder = append(mock.calls.GetLastStatusOrder, callInfo)
	mock.lockGetLastStatusOrder.Unlock()
	return mock.GetLastStatusOrderFunc(orderStatus, posApi)
}

// GetLastStatusOrderCalls gets all the calls that were made to GetLastStatusOrder.
// Check the length with:
//     len(mockedPosRepository.GetLastStatusOrderCalls())
func (mock *PosRepositoryMock) GetLastStatusOrderCalls() []struct {
	OrderStatus string
	PosApi      domain.PosResponseBodyApi
} {
	var calls []struct {
		OrderStatus string
		PosApi      domain.PosResponseBodyApi
	}
	mock.lockGetLastStatusOrder.RLock()
	calls = mock.calls.GetLastStatusOrder
	mock.lockGetLastStatusOrder.RUnlock()
	return calls
}

// Tracking calls TrackingFunc.
func (mock *PosRepositoryMock) Tracking(ctx context.Context, config domain.RepositoryParam) domain.ResultPosRepository {
	if mock.TrackingFunc == nil {
		panic("PosRepositoryMock.TrackingFunc: method is nil but PosRepository.Tracking was just called")
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
//     len(mockedPosRepository.TrackingCalls())
func (mock *PosRepositoryMock) TrackingCalls() []struct {
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