package service

import (
	"context"
	"errors"
	"testing"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/app/provider-jne/stub"
	"tracking-order-service/internal/pkg/constant"
)

func TestJneService_RunTracking(t *testing.T) {
	type fields struct {
		jneRepo JneRepository
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
				jneRepo: &JneRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultJneRepository {
						return domain.ResultJneRepository{
							Error: errors.New("error fetch data tracking from api"),
						}
					},
				},
			},
			args: args{},
			want: "GMs6XQBsuOM2",
		},
		{
			name: "case order tracking is empty",
			fields: fields{
				jneRepo: &JneRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultJneRepository {
						return domain.ResultJneRepository{
							ObjectResponse: domain.JneResponseBodyApi{
								History: []domain.JneResponseBodyApiHistory{
									{},
								},
							},
							RawResponse: stub.StubJneApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jneApi domain.JneResponseBodyApi) domain.ResultJneStatusOrder {
						return domain.ResultJneStatusOrder{
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
			want: "GMzOwFgFFknl",
		},
		{
			name: "case failed unmarshal db order tracking response",
			fields: fields{
				jneRepo: &JneRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultJneRepository {
						return domain.ResultJneRepository{
							ObjectResponse: domain.JneResponseBodyApi{
								History: []domain.JneResponseBodyApiHistory{
									{},
								},
							},
							RawResponse: stub.StubJneApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jneApi domain.JneResponseBodyApi) domain.ResultJneStatusOrder {
						return domain.ResultJneStatusOrder{
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
			want: "GMlGlTZjc6QX",
		},
		{
			name: "case db and api history length not match",
			fields: fields{
				jneRepo: &JneRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultJneRepository {
						return domain.ResultJneRepository{
							ObjectResponse: domain.JneResponseBodyApi{
								History: []domain.JneResponseBodyApiHistory{
									{},
								},
							},
							RawResponse: stub.StubJneApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jneApi domain.JneResponseBodyApi) domain.ResultJneStatusOrder {
						return domain.ResultJneStatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:  1,
					Response: stub.StubJneApi,
				},
			},
			want: "GMUDom1i573D",
		},
		{
			name: "case order not updated after max day",
			fields: fields{
				jneRepo: &JneRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultJneRepository {
						return domain.ResultJneRepository{
							ObjectResponse: domain.JneResponseBodyApi{
								History: []domain.JneResponseBodyApiHistory{
									{}, {}, {}, {}, {}, {}, {},
								},
							},
							RawResponse: stub.StubJneApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jneApi domain.JneResponseBodyApi) domain.ResultJneStatusOrder {
						return domain.ResultJneStatusOrder{
							Status: constant.SENDING,
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:    1,
					Response:   stub.StubJneApi,
					ModifiedAt: time.Now().AddDate(0, 0, -4),
				},
			},
			want: "GMJZHTLf97RC",
		},
		{
			name: "case no need to update",
			fields: fields{
				jneRepo: &JneRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultJneRepository {
						return domain.ResultJneRepository{
							ObjectResponse: domain.JneResponseBodyApi{
								History: []domain.JneResponseBodyApiHistory{
									{}, {}, {}, {}, {}, {}, {},
								},
							},
							RawResponse: stub.StubJneApi,
						}
					},
					GetLastStatusOrderFunc: func(orderStatus string, jneApi domain.JneResponseBodyApi) domain.ResultJneStatusOrder {
						return domain.ResultJneStatusOrder{
							Status:     constant.SENDING,
							ModifiedAt: time.Now(),
						}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:    1,
					Response:   stub.StubJneApi,
					ModifiedAt: time.Now(),
				},
			},
			want: "GMwuJeC7TRHn",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := NewJneService(tt.fields.jneRepo)
			if got := js.RunTracking(tt.args.order, &tt.args.rules); got.Hash != tt.want {
				t.Errorf("JneService.RunTracking() got = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}
