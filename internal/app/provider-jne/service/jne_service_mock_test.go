// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package service

import (
	"context"
	"sync"
	"tracking-order-service/internal/app/domain"
)

// Ensure, that JneRepositoryMock does implement JneRepository.
// If this is not the case, regenerate this file with moq.
var _ JneRepository = &JneRepositoryMock{}

// JneRepositoryMock is a mock implementation of JneRepository.
//
// 	func TestSomethingThatUsesJneRepository(t *testing.T) {
//
// 		// make and configure a mocked JneRepository
// 		mockedJneRepository := &JneRepositoryMock{
// 			GetLastStatusOrderFunc: func(orderStatus string, jneApi domain.JneResponseBodyApi) domain.ResultJneStatus {
// 				panic("mock out the GetLastStatusOrder method")
// 			},
// 			TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultJneRepository {
// 				panic("mock out the Tracking method")
// 			},
// 		}
//
// 		// use mockedJneRepository in code that requires JneRepository
// 		// and then make assertions.
//
// 	}
type JneRepositoryMock struct {
	// GetLastStatusOrderFunc mocks the GetLastStatusOrder method.
	GetLastStatusOrderFunc func(orderStatus string, jneApi domain.JneResponseBodyApi) domain.ResultJneStatusOrder

	// TrackingFunc mocks the Tracking method.
	TrackingFunc func(ctx context.Context, config domain.RepositoryParam) domain.ResultJneRepository

	// calls tracks calls to the methods.
	calls struct {
		// GetLastStatusOrder holds details about calls to the GetLastStatusOrder method.
		GetLastStatusOrder []struct {
			// OrderStatus is the orderStatus argument value.
			OrderStatus string
			// JneApi is the jneApi argument value.
			JneApi domain.JneResponseBodyApi
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
func (mock *JneRepositoryMock) GetLastStatusOrder(orderStatus string, jneApi domain.JneResponseBodyApi) domain.ResultJneStatusOrder {
	if mock.GetLastStatusOrderFunc == nil {
		panic("JneRepositoryMock.GetLastStatusOrderFunc: method is nil but JneRepository.GetLastStatusOrder was just called")
	}
	callInfo := struct {
		OrderStatus string
		JneApi      domain.JneResponseBodyApi
	}{
		OrderStatus: orderStatus,
		JneApi:      jneApi,
	}
	mock.lockGetLastStatusOrder.Lock()
	mock.calls.GetLastStatusOrder = append(mock.calls.GetLastStatusOrder, callInfo)
	mock.lockGetLastStatusOrder.Unlock()
	return mock.GetLastStatusOrderFunc(orderStatus, jneApi)
}

// GetLastStatusOrderCalls gets all the calls that were made to GetLastStatusOrder.
// Check the length with:
//     len(mockedJneRepository.GetLastStatusOrderCalls())
func (mock *JneRepositoryMock) GetLastStatusOrderCalls() []struct {
	OrderStatus string
	JneApi      domain.JneResponseBodyApi
} {
	var calls []struct {
		OrderStatus string
		JneApi      domain.JneResponseBodyApi
	}
	mock.lockGetLastStatusOrder.RLock()
	calls = mock.calls.GetLastStatusOrder
	mock.lockGetLastStatusOrder.RUnlock()
	return calls
}

// Tracking calls TrackingFunc.
func (mock *JneRepositoryMock) Tracking(ctx context.Context, config domain.RepositoryParam) domain.ResultJneRepository {
	if mock.TrackingFunc == nil {
		panic("JneRepositoryMock.TrackingFunc: method is nil but JneRepository.Tracking was just called")
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
//     len(mockedJneRepository.TrackingCalls())
func (mock *JneRepositoryMock) TrackingCalls() []struct {
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
