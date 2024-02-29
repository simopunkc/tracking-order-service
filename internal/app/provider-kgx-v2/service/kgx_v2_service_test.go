package service

import (
	"context"
	"errors"
	"testing"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/app/provider-kgx-v2/stub"
	"tracking-order-service/internal/pkg/constant"
)

func TestKgxV2Service_RunTracking(t *testing.T) {
	type fields struct {
		kgxV2Repo KgxV2Repository
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
				kgxV2Repo: &KgxV2RepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultKgxV2Repository {
						return domain.ResultKgxV2Repository{
							Error: errors.New("error fetch data tracking from api"),
						}
					},
				},
			},
			args: args{},
			want: "GMnMvql7oSKY",
		},
		{
			name: "case order tracking is empty",
			fields: fields{
				kgxV2Repo: &KgxV2RepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultKgxV2Repository {
						return domain.ResultKgxV2Repository{
							ObjectResponse: domain.KgxV2ResponseBodyApi{
								Data: []domain.KgxV2ResponseBodyApiData{
									{},
								},
							},
							RawResponse: stub.StubKgxV2Api,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, kgxV2Api domain.KgxV2ResponseBodyApi) domain.ResultKgxV2StatusOrder {
						return domain.ResultKgxV2StatusOrder{
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
			want: "GMoJpTgvbtk9",
		},
		{
			name: "case failed unmarshal db order tracking response",
			fields: fields{
				kgxV2Repo: &KgxV2RepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultKgxV2Repository {
						return domain.ResultKgxV2Repository{
							ObjectResponse: domain.KgxV2ResponseBodyApi{
								Data: []domain.KgxV2ResponseBodyApiData{
									{},
								},
							},
							RawResponse: stub.StubKgxV2Api,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, kgxV2Api domain.KgxV2ResponseBodyApi) domain.ResultKgxV2StatusOrder {
						return domain.ResultKgxV2StatusOrder{
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
			want: "GM5lVgaxG0yE",
		},
		{
			name: "case db and api history length not match",
			fields: fields{
				kgxV2Repo: &KgxV2RepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultKgxV2Repository {
						return domain.ResultKgxV2Repository{
							ObjectResponse: domain.KgxV2ResponseBodyApi{
								Data: []domain.KgxV2ResponseBodyApiData{
									{},
								},
							},
							RawResponse: stub.StubKgxV2Api,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, kgxV2Api domain.KgxV2ResponseBodyApi) domain.ResultKgxV2StatusOrder {
						return domain.ResultKgxV2StatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:  1,
					Response: stub.StubKgxV2Api,
				},
			},
			want: "GMhJNDAjEEBa",
		},
		{
			name: "case order not updated after max day",
			fields: fields{
				kgxV2Repo: &KgxV2RepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultKgxV2Repository {
						return domain.ResultKgxV2Repository{
							ObjectResponse: domain.KgxV2ResponseBodyApi{
								Data: []domain.KgxV2ResponseBodyApiData{
									{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
								},
							},
							RawResponse: stub.StubKgxV2Api,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, kgxV2Api domain.KgxV2ResponseBodyApi) domain.ResultKgxV2StatusOrder {
						return domain.ResultKgxV2StatusOrder{
							Status: constant.SENDING,
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:    1,
					Response:   stub.StubKgxV2Api,
					ModifiedAt: time.Now().AddDate(0, 0, -4),
				},
			},
			want: "GM8e1M64zsQu",
		},
		{
			name: "case no need to update",
			fields: fields{
				kgxV2Repo: &KgxV2RepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultKgxV2Repository {
						return domain.ResultKgxV2Repository{
							ObjectResponse: domain.KgxV2ResponseBodyApi{
								Data: []domain.KgxV2ResponseBodyApiData{
									{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
								},
							},
							RawResponse: stub.StubKgxV2Api,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, kgxV2Api domain.KgxV2ResponseBodyApi) domain.ResultKgxV2StatusOrder {
						return domain.ResultKgxV2StatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:    1,
					Response:   stub.StubKgxV2Api,
					ModifiedAt: time.Now(),
				},
			},
			want: "GMrkgsMZmRGp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kvs := NewKgxV2Service(tt.fields.kgxV2Repo)
			if got := kvs.RunTracking(tt.args.order, &tt.args.rules); got.Hash != tt.want {
				t.Errorf("KgxV2Service.RunTracking() got = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}
