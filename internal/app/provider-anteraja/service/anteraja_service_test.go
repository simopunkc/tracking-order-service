package service

import (
	"context"
	"errors"
	"testing"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/app/provider-anteraja/stub"
	"tracking-order-service/internal/pkg/constant"
)

func TestAnterajaService_RunTracking(t *testing.T) {
	type fields struct {
		anterajaRepo AnterajaRepository
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
				anterajaRepo: &AnterajaRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultAnterajaRepository {
						return domain.ResultAnterajaRepository{
							Error: errors.New("error fetch data tracking from api"),
						}
					},
				},
			},
			args: args{},
			want: "GMBEJkbD6SPm",
		},
		{
			name: "case not 200 ok",
			fields: fields{
				anterajaRepo: &AnterajaRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultAnterajaRepository {
						return domain.ResultAnterajaRepository{
							ObjectResponse: domain.AnterajaResponseBodyApi{
								Status: 400,
							},
							RawResponse: stub.StubAnterajaApiNotFound,
						}
					},
				},
			},
			args: args{},
			want: "GMT4bPm9ecBA",
		},
		{
			name: "case order tracking is empty",
			fields: fields{
				anterajaRepo: &AnterajaRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultAnterajaRepository {
						return domain.ResultAnterajaRepository{
							ObjectResponse: domain.AnterajaResponseBodyApi{
								Status: 200,
								Content: domain.AnterajaResponseBodyApiContent{
									History: []domain.AnterajaResponseBodyApiContentHistory{
										{},
									},
								},
							},
							RawResponse: stub.StubAnterajaApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jntApi domain.AnterajaResponseBodyApi) domain.ResultAnterajaStatusOrder {
						return domain.ResultAnterajaStatusOrder{
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
			want: "GMv19TqabRlJ",
		},
		{
			name: "case failed unmarshal db order tracking response",
			fields: fields{
				anterajaRepo: &AnterajaRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultAnterajaRepository {
						return domain.ResultAnterajaRepository{
							ObjectResponse: domain.AnterajaResponseBodyApi{
								Status: 200,
								Content: domain.AnterajaResponseBodyApiContent{
									History: []domain.AnterajaResponseBodyApiContentHistory{
										{},
									},
								},
							},
							RawResponse: stub.StubAnterajaApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jntApi domain.AnterajaResponseBodyApi) domain.ResultAnterajaStatusOrder {
						return domain.ResultAnterajaStatusOrder{
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
			want: "GMWpVLpeTJgG",
		},
		{
			name: "case db and api history length not match",
			fields: fields{
				anterajaRepo: &AnterajaRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultAnterajaRepository {
						return domain.ResultAnterajaRepository{
							ObjectResponse: domain.AnterajaResponseBodyApi{
								Status: 200,
								Content: domain.AnterajaResponseBodyApiContent{
									History: []domain.AnterajaResponseBodyApiContentHistory{
										{},
									},
								},
							},
							RawResponse: stub.StubAnterajaApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jntApi domain.AnterajaResponseBodyApi) domain.ResultAnterajaStatusOrder {
						return domain.ResultAnterajaStatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:  1,
					Response: stub.StubAnterajaApi,
				},
			},
			want: "GMXWyuXNnMGg",
		},
		{
			name: "case order not updated after max day",
			fields: fields{
				anterajaRepo: &AnterajaRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultAnterajaRepository {
						return domain.ResultAnterajaRepository{
							ObjectResponse: domain.AnterajaResponseBodyApi{
								Status: 200,
								Content: domain.AnterajaResponseBodyApiContent{
									History: []domain.AnterajaResponseBodyApiContentHistory{
										{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
									},
								},
							},
							RawResponse: stub.StubAnterajaApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jntApi domain.AnterajaResponseBodyApi) domain.ResultAnterajaStatusOrder {
						return domain.ResultAnterajaStatusOrder{
							Status: constant.SENDING,
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:    1,
					Response:   stub.StubAnterajaApi,
					ModifiedAt: time.Now().AddDate(0, 0, -4),
				},
			},
			want: "GM8Ebv74zWl2",
		},
		{
			name: "case no need to update",
			fields: fields{
				anterajaRepo: &AnterajaRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultAnterajaRepository {
						return domain.ResultAnterajaRepository{
							ObjectResponse: domain.AnterajaResponseBodyApi{
								Status: 200,
								Content: domain.AnterajaResponseBodyApiContent{
									History: []domain.AnterajaResponseBodyApiContentHistory{
										{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
									},
								},
							},
							RawResponse: stub.StubAnterajaApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jntApi domain.AnterajaResponseBodyApi) domain.ResultAnterajaStatusOrder {
						return domain.ResultAnterajaStatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:    1,
					Response:   stub.StubAnterajaApi,
					ModifiedAt: time.Now(),
				},
			},
			want: "GMSjB3s2Dvot",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := NewAnterajaService(tt.fields.anterajaRepo)
			if got := as.RunTracking(tt.args.order, &tt.args.rules); got.Hash != tt.want {
				t.Errorf("AnterajaService.RunTracking() = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}
