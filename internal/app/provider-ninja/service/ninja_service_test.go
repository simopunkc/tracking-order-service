package service

import (
	"context"
	"errors"
	"testing"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/app/provider-ninja/stub"
)

func TestNinjaService_RunTracking(t *testing.T) {
	type fields struct {
		ninjaRepo NinjaRepository
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
				ninjaRepo: &NinjaRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultNinjaRepository {
						return domain.ResultNinjaRepository{
							Error: errors.New("error fetch data tracking from api"),
						}
					},
				},
			},
			args: args{},
			want: "GMxsmc9q1het",
		},
		{
			name: "case no need to update",
			fields: fields{
				ninjaRepo: &NinjaRepositoryMock{
					TrackingFunc: func(ctx context.Context, config domain.RepositoryParam) domain.ResultNinjaRepository {
						return domain.ResultNinjaRepository{
							ObjectResponse: domain.NinjaResponseBodyApi{
								ShipperID: 1,
							},
						}
					},
					GetLastStatusOrderFunc: func(orderTracking []domain.NinjaResponseBodyApi, ninjaApi domain.NinjaResponseBodyApi) domain.ListNinjaResponseBodyApi {
						return domain.ListNinjaResponseBodyApi{}
					},
				},
			},
			args: args{
				order: domain.OrderTracking{
					OrderID:    1,
					Response:   stub.StubNinjaDatabase,
					ModifiedAt: time.Now(),
				},
			},
			want: "GMdiUEk2E7R8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := NewNinjaService(tt.fields.ninjaRepo)
			if got := ns.RunTracking(tt.args.order, &tt.args.rules); got.Hash != tt.want {
				t.Errorf("NinjaService.RunTracking() got = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}
