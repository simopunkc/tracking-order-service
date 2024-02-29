package service

import (
	"context"
	"errors"
	"testing"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/app/provider-jnt/stub"
	"tracking-order-service/internal/pkg/constant"
)

func TestJntService_RunTracking(t *testing.T) {
	type fields struct {
		jntRepo JntRepository
	}
	type args struct {
		order domain.OrderTracking
		rules domain.SettingsValueCheckOrder
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "case failed retrieve tracking",
			fields: fields{
				jntRepo: &JntRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultJntRepository {
						return domain.ResultJntRepository{
							Error: errors.New("error fetch data tracking from api"),
						}
					},
				},
			},
			args: args{},
			want: "GMOa8O8UJskH",
		},
		{
			name: "case order tracking is empty",
			fields: fields{
				jntRepo: &JntRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultJntRepository {
						return domain.ResultJntRepository{
							ObjectResponse: domain.JntResponseBodyApi{
								History: []domain.JntResponseBodyApiHistory{
									{},
								},
							},
							RawResponse: stub.StubJntApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jntApi domain.JntResponseBodyApi) domain.ResultJntStatusOrder {
						return domain.ResultJntStatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID: 0,
				},
			},
			want: "GM2zqs0JxDAk",
		},
		{
			name: "case failed unmarshal db order tracking response",
			fields: fields{
				jntRepo: &JntRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultJntRepository {
						return domain.ResultJntRepository{
							ObjectResponse: domain.JntResponseBodyApi{
								History: []domain.JntResponseBodyApiHistory{
									{},
								},
							},
							RawResponse: stub.StubJntApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jntApi domain.JntResponseBodyApi) domain.ResultJntStatusOrder {
						return domain.ResultJntStatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:  1,
					Response: constant.EmptyString,
				},
			},
			want: "GMZ2HpjpIK8K",
		},
		{
			name: "case db and api history length not match",
			fields: fields{
				jntRepo: &JntRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultJntRepository {
						return domain.ResultJntRepository{
							ObjectResponse: domain.JntResponseBodyApi{
								History: []domain.JntResponseBodyApiHistory{
									{},
								},
							},
							RawResponse: stub.StubJntApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jntApi domain.JntResponseBodyApi) domain.ResultJntStatusOrder {
						return domain.ResultJntStatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:  1,
					Response: stub.StubJntApi,
				},
			},
			want: "GMLng7IsolKR",
		},
		{
			name: "case order not updated after max day",
			fields: fields{
				jntRepo: &JntRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultJntRepository {
						return domain.ResultJntRepository{
							ObjectResponse: domain.JntResponseBodyApi{
								History: []domain.JntResponseBodyApiHistory{
									{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
								},
							},
							RawResponse: stub.StubJntApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jntApi domain.JntResponseBodyApi) domain.ResultJntStatusOrder {
						return domain.ResultJntStatusOrder{
							Status: constant.SENDING,
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:    1,
					Response:   stub.StubJntApi,
					ModifiedAt: time.Now().AddDate(0, 0, -4),
				},
			},
			want: "GMquingxcWCe",
		},
		{
			name: "case no need to update",
			fields: fields{
				jntRepo: &JntRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultJntRepository {
						return domain.ResultJntRepository{
							ObjectResponse: domain.JntResponseBodyApi{
								History: []domain.JntResponseBodyApiHistory{
									{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
								},
							},
							RawResponse: stub.StubJntApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jntApi domain.JntResponseBodyApi) domain.ResultJntStatusOrder {
						return domain.ResultJntStatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:    1,
					Response:   stub.StubJntApi,
					ModifiedAt: time.Now(),
				},
			},
			want: "GMtlsUlxC7SI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := NewJntService(tt.fields.jntRepo)
			if got := js.RunTracking(tt.args.order, &tt.args.rules); got.Hash != tt.want {
				t.Errorf("JntService.RunTracking() got = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}
