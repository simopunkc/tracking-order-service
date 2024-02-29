package service

import (
	"context"
	"errors"
	"testing"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/app/provider-sicepat/stub"
	"tracking-order-service/internal/pkg/constant"
)

func TestSicepatService_RunTracking(t *testing.T) {
	type fields struct {
		sicepatRepo SicepatRepository
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
				sicepatRepo: &SicepatRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultSicepatRepository {
						return domain.ResultSicepatRepository{
							Error: errors.New("error fetch data tracking from api"),
						}
					},
				},
			},
			args: args{},
			want: "GMmNattIDtUg",
		},
		{
			name: "case not 200 ok",
			fields: fields{
				sicepatRepo: &SicepatRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultSicepatRepository {
						return domain.ResultSicepatRepository{
							ObjectResponse: domain.SicepatResponseBodyApi{
								Sicepat: domain.SicepatResponseBodyApiSicepat{
									Status: domain.SicepatResponseBodyApiSicepatStatus{
										Code: 400,
									},
								},
							},
							RawResponse: stub.MockSicepatApiNotFound,
						}
					},
				},
			},
			args: args{},
			want: "GM9z3BeFFaUT",
		},
		{
			name: "case order tracking is empty",
			fields: fields{
				sicepatRepo: &SicepatRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultSicepatRepository {
						return domain.ResultSicepatRepository{
							ObjectResponse: domain.SicepatResponseBodyApi{
								Sicepat: domain.SicepatResponseBodyApiSicepat{
									Status: domain.SicepatResponseBodyApiSicepatStatus{
										Code: 200,
									},
									Result: domain.SicepatResponseBodyApiSicepatResult{
										TrackHistory: []domain.SicepatResponseBodyApiSicepatResultTrackHistory{
											{},
										},
									},
								},
							},
							RawResponse: stub.MockSicepatApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, sicepatApi domain.SicepatResponseBodyApi) domain.ResultSicepatStatusOrder {
						return domain.ResultSicepatStatusOrder{
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
			want: "GMvzuHX2PKva",
		},
		{
			name: "case failed unmarshal db order tracking response",
			fields: fields{
				sicepatRepo: &SicepatRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultSicepatRepository {
						return domain.ResultSicepatRepository{
							ObjectResponse: domain.SicepatResponseBodyApi{
								Sicepat: domain.SicepatResponseBodyApiSicepat{
									Status: domain.SicepatResponseBodyApiSicepatStatus{
										Code: 200,
									},
									Result: domain.SicepatResponseBodyApiSicepatResult{
										TrackHistory: []domain.SicepatResponseBodyApiSicepatResultTrackHistory{
											{},
										},
									},
								},
							},
							RawResponse: stub.MockSicepatApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, sicepatApi domain.SicepatResponseBodyApi) domain.ResultSicepatStatusOrder {
						return domain.ResultSicepatStatusOrder{
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
			want: "GMw6XJsF5TiU",
		},
		{
			name: "case db and api history length not match",
			fields: fields{
				sicepatRepo: &SicepatRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultSicepatRepository {
						return domain.ResultSicepatRepository{
							ObjectResponse: domain.SicepatResponseBodyApi{
								Sicepat: domain.SicepatResponseBodyApiSicepat{
									Status: domain.SicepatResponseBodyApiSicepatStatus{
										Code: 200,
									},
									Result: domain.SicepatResponseBodyApiSicepatResult{
										TrackHistory: []domain.SicepatResponseBodyApiSicepatResultTrackHistory{
											{},
										},
									},
								},
							},
							RawResponse: stub.MockSicepatApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, sicepatApi domain.SicepatResponseBodyApi) domain.ResultSicepatStatusOrder {
						return domain.ResultSicepatStatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:  1,
					Response: stub.MockSicepatApi,
				},
			},
			want: "GMsSBuMBnzuG",
		},
		{
			name: "case unmatch last time date time",
			fields: fields{
				sicepatRepo: &SicepatRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultSicepatRepository {
						return domain.ResultSicepatRepository{
							ObjectResponse: domain.SicepatResponseBodyApi{
								Sicepat: domain.SicepatResponseBodyApiSicepat{
									Status: domain.SicepatResponseBodyApiSicepatStatus{
										Code: 200,
									},
									Result: domain.SicepatResponseBodyApiSicepatResult{
										LastStatus: domain.SicepatResponseBodyApiSicepatResultLastStatus{
											DateTime: constant.EmptyString,
										},
										TrackHistory: []domain.SicepatResponseBodyApiSicepatResultTrackHistory{
											{}, {}, {}, {}, {},
										},
									},
								},
							},
							RawResponse: stub.MockSicepatApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, sicepatApi domain.SicepatResponseBodyApi) domain.ResultSicepatStatusOrder {
						return domain.ResultSicepatStatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:  1,
					Response: stub.MockSicepatApi,
				},
			},
			want: "GMvM0YohP1lu",
		},
		{
			name: "case order not updated after max day",
			fields: fields{
				sicepatRepo: &SicepatRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultSicepatRepository {
						return domain.ResultSicepatRepository{
							ObjectResponse: domain.SicepatResponseBodyApi{
								Sicepat: domain.SicepatResponseBodyApiSicepat{
									Status: domain.SicepatResponseBodyApiSicepatStatus{
										Code: 200,
									},
									Result: domain.SicepatResponseBodyApiSicepatResult{
										LastStatus: domain.SicepatResponseBodyApiSicepatResultLastStatus{
											DateTime: stub.MockSicepatApiLastUpdated,
										},
										TrackHistory: []domain.SicepatResponseBodyApiSicepatResultTrackHistory{
											{}, {}, {}, {}, {},
										},
									},
								},
							},
							RawResponse: stub.MockSicepatApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, sicepatApi domain.SicepatResponseBodyApi) domain.ResultSicepatStatusOrder {
						return domain.ResultSicepatStatusOrder{
							Status: constant.SENDING,
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:    1,
					Response:   stub.MockSicepatApi,
					ModifiedAt: time.Now().AddDate(0, 0, -4),
				},
			},
			want: "GMGQ1HTnw2vs",
		},
		{
			name: "case no need to update",
			fields: fields{
				sicepatRepo: &SicepatRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultSicepatRepository {
						return domain.ResultSicepatRepository{
							ObjectResponse: domain.SicepatResponseBodyApi{
								Sicepat: domain.SicepatResponseBodyApiSicepat{
									Status: domain.SicepatResponseBodyApiSicepatStatus{
										Code: 200,
									},
									Result: domain.SicepatResponseBodyApiSicepatResult{
										LastStatus: domain.SicepatResponseBodyApiSicepatResultLastStatus{
											DateTime: stub.MockSicepatApiLastUpdated,
										},
										TrackHistory: []domain.SicepatResponseBodyApiSicepatResultTrackHistory{
											{}, {}, {}, {}, {},
										},
									},
								},
							},
							RawResponse: stub.MockSicepatApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, sicepatApi domain.SicepatResponseBodyApi) domain.ResultSicepatStatusOrder {
						return domain.ResultSicepatStatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:    1,
					Response:   stub.MockSicepatApi,
					ModifiedAt: time.Now(),
				},
			},
			want: "GM5j2RzZRraO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := NewSicepatService(tt.fields.sicepatRepo)
			if got := ss.RunTracking(tt.args.order, &tt.args.rules); got.Hash != tt.want {
				t.Errorf("SicepatService.RunTracking() got = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}
