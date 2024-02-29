package service

import (
	"context"
	"errors"
	"testing"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/app/provider-pos/stub"
	"tracking-order-service/internal/pkg/constant"
)

func TestPosService_RunTracking(t *testing.T) {
	type fields struct {
		posRepo PosRepository
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
				posRepo: &PosRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultPosRepository {
						return domain.ResultPosRepository{
							Error: errors.New("error fetch data tracking from api"),
						}
					},
				},
			},
			args: args{},
			want: "GMVdeWwgd80S",
		},
		{
			name: "case order tracking is empty",
			fields: fields{
				posRepo: &PosRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultPosRepository {
						return domain.ResultPosRepository{
							ObjectResponse: domain.PosResponseBodyApi{
								Response: domain.PosResponseBodyApiResponse{
									Data: []domain.PosResponseBodyApiResponseData{
										{},
									},
								},
							},
							RawResponse: stub.StubPosApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, posApi domain.PosResponseBodyApi) domain.ResultPosStatusOrder {
						return domain.ResultPosStatusOrder{
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
			want: "GMlIEjkOokSJ",
		},
		{
			name: "case failed unmarshal db order tracking response",
			fields: fields{
				posRepo: &PosRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultPosRepository {
						return domain.ResultPosRepository{
							ObjectResponse: domain.PosResponseBodyApi{
								Response: domain.PosResponseBodyApiResponse{
									Data: []domain.PosResponseBodyApiResponseData{
										{},
									},
								},
							},
							RawResponse: stub.StubPosApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, posApi domain.PosResponseBodyApi) domain.ResultPosStatusOrder {
						return domain.ResultPosStatusOrder{
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
			want: "GMyO1jQpZfh8",
		},
		{
			name: "case db and api history length not match",
			fields: fields{
				posRepo: &PosRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultPosRepository {
						return domain.ResultPosRepository{
							ObjectResponse: domain.PosResponseBodyApi{
								Response: domain.PosResponseBodyApiResponse{
									Data: []domain.PosResponseBodyApiResponseData{
										{},
									},
								},
							},
							RawResponse: stub.StubPosApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, posApi domain.PosResponseBodyApi) domain.ResultPosStatusOrder {
						return domain.ResultPosStatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:  1,
					Response: stub.StubPosApi,
				},
			},
			want: "GMyvln2PQbxV",
		},
		{
			name: "case order not updated after max day",
			fields: fields{
				posRepo: &PosRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultPosRepository {
						return domain.ResultPosRepository{
							ObjectResponse: domain.PosResponseBodyApi{
								Response: domain.PosResponseBodyApiResponse{
									Data: []domain.PosResponseBodyApiResponseData{
										{}, {}, {}, {}, {},
									},
								},
							},
							RawResponse: stub.StubPosApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, posApi domain.PosResponseBodyApi) domain.ResultPosStatusOrder {
						return domain.ResultPosStatusOrder{
							Status: constant.SENDING,
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:    1,
					Response:   stub.StubPosApi,
					ModifiedAt: time.Now().AddDate(0, 0, -4),
				},
			},
			want: "GMjdDjebsCox",
		},
		{
			name: "case no need to update",
			fields: fields{
				posRepo: &PosRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultPosRepository {
						return domain.ResultPosRepository{
							ObjectResponse: domain.PosResponseBodyApi{
								Response: domain.PosResponseBodyApiResponse{
									Data: []domain.PosResponseBodyApiResponseData{
										{}, {}, {}, {}, {},
									},
								},
							},
							RawResponse: stub.StubPosApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, posApi domain.PosResponseBodyApi) domain.ResultPosStatusOrder {
						return domain.ResultPosStatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:    1,
					Response:   stub.StubPosApi,
					ModifiedAt: time.Now(),
				},
			},
			want: "GM8dDIHq3aUw",
		},
		{
			name: "case no need to update, using api staging",
			fields: fields{
				posRepo: &PosRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultPosRepository {
						return domain.ResultPosRepository{
							ObjectResponse: domain.PosResponseBodyApi{
								Response: domain.PosResponseBodyApiResponse{
									Data: []domain.PosResponseBodyApiResponseData{
										{}, {}, {}, {}, {},
									},
								},
							},
							RawResponse: stub.StubPosApiStagingNotFound,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, posApi domain.PosResponseBodyApi) domain.ResultPosStatusOrder {
						return domain.ResultPosStatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:    1,
					Response:   stub.StubPosApiStaging,
					ModifiedAt: time.Now(),
				},
			},
			want: "GM8dDIHq3aUw",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := NewPosService(tt.fields.posRepo)
			if got := ps.RunTracking(tt.args.order, &tt.args.rules); got.Hash != tt.want {
				t.Errorf("PosService.RunTracking() got = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}
