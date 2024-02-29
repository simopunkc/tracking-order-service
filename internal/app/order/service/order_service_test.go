package service

import (
	"errors"
	"testing"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/app/order/config"
	"tracking-order-service/internal/app/order/stub"
	"tracking-order-service/internal/pkg/constant"
)

func TestOrderService_RunCronJob(t *testing.T) {
	type fields struct {
		orderRepoRead  DatabaseReadOrderRepository
		orderRepoWrite DatabaseWriteOrderRepository
		anterajaServ   ProviderService
		jneServ        ProviderService
		jntServ        ProviderService
		kgxv2Serv      ProviderService
		ninjaServ      ProviderService
		posServ        ProviderService
		sicepatServ    ProviderService
	}
	type args struct {
		argument domain.TrackingOrderParam
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "case orders model error",
			fields: fields{
				orderRepoRead: &DatabaseReadOrderRepositoryMock{
					OrdersModelFunc: func(config domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
						return []domain.OrderTracking{}, errors.New("error fetch order from db")
					},
				},
			},
			args: args{},
			want: "GMtJXpTDjO7R",
		},
		{
			name: "case orders model length is 0",
			fields: fields{
				orderRepoRead: &DatabaseReadOrderRepositoryMock{
					OrdersModelFunc: func(config domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
						return []domain.OrderTracking{}, nil
					},
				},
			},
			args: args{},
			want: "GMi7Y67vyz6B",
		},
		{
			name: "case rules model is error",
			fields: fields{
				orderRepoRead: &DatabaseReadOrderRepositoryMock{
					OrdersModelFunc: func(config domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
						return []domain.OrderTracking{
							{
								ID: 1,
							},
							{
								ID: 2,
							},
						}, nil
					},
					RulesModelFunc: func(config domain.TrackingOrderParam) (domain.Settings, error) {
						return domain.Settings{}, errors.New("error fetch rules from db")
					},
				},
			},
			args: args{},
			want: "GMQ5qIDozYcI",
		},
		{
			name: "case unmarshal rules model error",
			fields: fields{
				orderRepoRead: &DatabaseReadOrderRepositoryMock{
					OrdersModelFunc: func(config domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
						return []domain.OrderTracking{
							{
								ID: 1,
							},
							{
								ID: 2,
							},
						}, nil
					},
					RulesModelFunc: func(config domain.TrackingOrderParam) (domain.Settings, error) {
						return domain.Settings{
							Value: constant.EmptyString,
						}, nil
					},
				},
			},
			args: args{},
			want: "GM5GTbniGiv0",
		},
		{
			name: "case failed update 1 order tracking",
			fields: fields{
				orderRepoRead: &DatabaseReadOrderRepositoryMock{
					OrdersModelFunc: func(config domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
						return []domain.OrderTracking{
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.ANTERAJA,
									},
								},
							},
						}, nil
					},
					RulesModelFunc: func(config domain.TrackingOrderParam) (domain.Settings, error) {
						return domain.Settings{
							Value: stub.StubDbSettingsCheckOrderValue,
						}, nil
					},
				},
				orderRepoWrite: &DatabaseWriteOrderRepositoryMock{
					BulkUpdateOrderTrackingFunc: func(config []domain.OrderTracking) (int64, error) {
						return 0, errors.New("error update order tracking")
					},
				},
				anterajaServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				jneServ:     &ProviderServiceMock{},
				jntServ:     &ProviderServiceMock{},
				kgxv2Serv:   &ProviderServiceMock{},
				ninjaServ:   &ProviderServiceMock{},
				posServ:     &ProviderServiceMock{},
				sicepatServ: &ProviderServiceMock{},
			},
			args: args{},
			want: "GMasJ1wQqVZJ",
		},
		{
			name: "case success update 7 order tracking",
			fields: fields{
				orderRepoRead: &DatabaseReadOrderRepositoryMock{
					OrdersModelFunc: func(config domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
						return []domain.OrderTracking{
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.SICEPAT,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.ANTERAJA,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.KGX,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.JNT,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.JNE,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.POS,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.NINJA,
									},
								},
							},
						}, nil
					},
					RulesModelFunc: func(config domain.TrackingOrderParam) (domain.Settings, error) {
						return domain.Settings{
							Value: stub.StubDbSettingsCheckOrderValue,
						}, nil
					},
				},
				orderRepoWrite: &DatabaseWriteOrderRepositoryMock{
					BulkUpdateOrderTrackingFunc: func(config []domain.OrderTracking) (int64, error) {
						return 6, nil
					},
				},
				anterajaServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				jneServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				jntServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				kgxv2Serv: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				ninjaServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				posServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				sicepatServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								ModifiedAt: time.Now(),
							},
						}
					},
				},
			},
			args: args{},
			want: "GMWg8IfAyCHH",
		},
		{
			name: "case failed update 1 order",
			fields: fields{
				orderRepoRead: &DatabaseReadOrderRepositoryMock{
					OrdersModelFunc: func(config domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
						return []domain.OrderTracking{
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.SICEPAT,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.ANTERAJA,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.KGX,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.JNT,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.JNE,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.POS,
									},
								},
							},
						}, nil
					},
					RulesModelFunc: func(config domain.TrackingOrderParam) (domain.Settings, error) {
						return domain.Settings{
							Value: stub.StubDbSettingsCheckOrderValue,
						}, nil
					},
				},
				orderRepoWrite: &DatabaseWriteOrderRepositoryMock{
					BulkUpdateOrderTrackingFunc: func(config []domain.OrderTracking) (int64, error) {
						return 1, nil
					},
					BulkUpdateOrderFunc: func(config []domain.Order) (int64, error) {
						return 0, errors.New("error update order")
					},
				},
				anterajaServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								Order: domain.Order{
									Status:     constant.DELIVERED,
									ModifiedAt: time.Now().AddDate(0, 0, 1),
								},
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				jneServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
				jntServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
				kgxv2Serv: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
				ninjaServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				posServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
				sicepatServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
			},
			args: args{},
			want: "GMmPd7S1SVRB",
		},
		{
			name: "case success update 12 object",
			fields: fields{
				orderRepoRead: &DatabaseReadOrderRepositoryMock{
					OrdersModelFunc: func(config domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
						return []domain.OrderTracking{
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.SICEPAT,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.ANTERAJA,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.KGX,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.JNT,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.JNE,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.POS,
									},
								},
							},
						}, nil
					},
					RulesModelFunc: func(config domain.TrackingOrderParam) (domain.Settings, error) {
						return domain.Settings{
							Value: stub.StubDbSettingsCheckOrderValue,
						}, nil
					},
				},
				orderRepoWrite: &DatabaseWriteOrderRepositoryMock{
					BulkUpdateOrderTrackingFunc: func(config []domain.OrderTracking) (int64, error) {
						return 6, nil
					},
					BulkUpdateOrderFunc: func(config []domain.Order) (int64, error) {
						return 6, nil
					},
				},
				anterajaServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								Order: domain.Order{
									Status:     constant.DELIVERED,
									ModifiedAt: time.Now().AddDate(0, 0, 1),
								},
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				jneServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								Order: domain.Order{
									Status:     constant.DELIVERED,
									ModifiedAt: time.Now().AddDate(0, 0, 1),
								},
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				jntServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								Order: domain.Order{
									Status:     constant.DELIVERED,
									ModifiedAt: time.Now().AddDate(0, 0, 1),
								},
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				kgxv2Serv: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								Order: domain.Order{
									Status:     constant.DELIVERED,
									ModifiedAt: time.Now().AddDate(0, 0, 1),
								},
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				ninjaServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				posServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								Order: domain.Order{
									Status:     constant.DELIVERED,
									ModifiedAt: time.Now().AddDate(0, 0, 1),
								},
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				sicepatServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								Order: domain.Order{
									Status:     constant.DELIVERED,
									ModifiedAt: time.Now().AddDate(0, 0, 1),
								},
								ModifiedAt: time.Now(),
							},
						}
					},
				},
			},
			args: args{},
			want: "GMWg8IfAyCHH",
		},
		{
			name: "case no updated order tracking",
			fields: fields{
				orderRepoRead: &DatabaseReadOrderRepositoryMock{
					OrdersModelFunc: func(config domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
						return []domain.OrderTracking{
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.SICEPAT,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.ANTERAJA,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.KGX,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.JNT,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.JNE,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.POS,
									},
								},
							},
						}, nil
					},
					RulesModelFunc: func(config domain.TrackingOrderParam) (domain.Settings, error) {
						return domain.Settings{
							Value: stub.StubDbSettingsCheckOrderValue,
						}, nil
					},
				},
				anterajaServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
				jneServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
				jntServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
				kgxv2Serv: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
				ninjaServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				posServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
				sicepatServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
			},
			args: args{},
			want: "GMWg8IfAyCHH",
		},
		{
			name: "case unknown order tracking provider slug",
			fields: fields{
				orderRepoRead: &DatabaseReadOrderRepositoryMock{
					OrdersModelFunc: func(config domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
						return []domain.OrderTracking{
							{
								Order: domain.Order{
									Provider: domain.Provider{},
								},
							},
						}, nil
					},
					RulesModelFunc: func(config domain.TrackingOrderParam) (domain.Settings, error) {
						return domain.Settings{
							Value: stub.StubDbSettingsCheckOrderValue,
						}, nil
					},
				},
				orderRepoWrite: &DatabaseWriteOrderRepositoryMock{},
				anterajaServ:   &ProviderServiceMock{},
				jneServ:        &ProviderServiceMock{},
				jntServ:        &ProviderServiceMock{},
				kgxv2Serv:      &ProviderServiceMock{},
				ninjaServ:      &ProviderServiceMock{},
				posServ:        &ProviderServiceMock{},
				sicepatServ:    &ProviderServiceMock{},
			},
			args: args{},
			want: "GMWg8IfAyCHH",
		},
		{
			name: "case deadline runtask exceeded",
			fields: fields{
				orderRepoRead: &DatabaseReadOrderRepositoryMock{
					OrdersModelFunc: func(config domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
						return []domain.OrderTracking{
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.SICEPAT,
									},
								},
							},
						}, nil
					},
					RulesModelFunc: func(config domain.TrackingOrderParam) (domain.Settings, error) {
						return domain.Settings{
							Value: stub.StubDbSettingsCheckOrderValue,
						}, nil
					},
				},
				orderRepoWrite: &DatabaseWriteOrderRepositoryMock{},
				anterajaServ:   &ProviderServiceMock{},
				jneServ:        &ProviderServiceMock{},
				jntServ:        &ProviderServiceMock{},
				kgxv2Serv:      &ProviderServiceMock{},
				ninjaServ:      &ProviderServiceMock{},
				posServ:        &ProviderServiceMock{},
				sicepatServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						const deadlineRunTask = config.MaxSecondDeadlineRunTask + 1
						time.Sleep(deadlineRunTask * time.Second)
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
			},
			args: args{},
			want: "GMWg8IfAyCHH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os := NewOrderService(tt.fields.orderRepoRead, tt.fields.orderRepoWrite, tt.fields.anterajaServ, tt.fields.jneServ, tt.fields.jntServ, tt.fields.kgxv2Serv, tt.fields.ninjaServ, tt.fields.posServ, tt.fields.sicepatServ)
			if got := os.RunCronJob(tt.args.argument); got.Hash != tt.want {
				t.Errorf("OrderService.RunCronJob() = %v, want %v", got.Hash, tt.want)
			}
		})
	}
}

func BenchmarkOrderService_RunCronJob(b *testing.B) {
	type fields struct {
		orderRepoRead  DatabaseReadOrderRepository
		orderRepoWrite DatabaseWriteOrderRepository
		anterajaServ   ProviderService
		jneServ        ProviderService
		jntServ        ProviderService
		kgxv2Serv      ProviderService
		ninjaServ      ProviderService
		posServ        ProviderService
		sicepatServ    ProviderService
	}
	type args struct {
		argument domain.TrackingOrderParam
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "case no updated order tracking",
			fields: fields{
				orderRepoRead: &DatabaseReadOrderRepositoryMock{
					OrdersModelFunc: func(config domain.TrackingOrderParam) ([]domain.OrderTracking, error) {
						return []domain.OrderTracking{
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.SICEPAT,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.ANTERAJA,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.KGX,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.JNT,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.JNE,
									},
								},
							},
							{
								Order: domain.Order{
									Provider: domain.Provider{
										Slug: constant.POS,
									},
								},
							},
						}, nil
					},
					RulesModelFunc: func(config domain.TrackingOrderParam) (domain.Settings, error) {
						return domain.Settings{
							Value: stub.StubDbSettingsCheckOrderValue,
						}, nil
					},
				},
				anterajaServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
				jneServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
				jntServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
				kgxv2Serv: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
				ninjaServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{
								ModifiedAt: time.Now(),
							},
						}
					},
				},
				posServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
				sicepatServ: &ProviderServiceMock{
					RunTrackingFunc: func(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
						return domain.ResultProviderService{
							NewOrder: domain.OrderTracking{},
						}
					},
				},
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		for i := 0; i < b.N; i++ {
			os := OrderService{
				orderRepoRead:  tt.fields.orderRepoRead,
				orderRepoWrite: tt.fields.orderRepoWrite,
				anterajaServ:   tt.fields.anterajaServ,
				jneServ:        tt.fields.jneServ,
				jntServ:        tt.fields.jntServ,
				kgxv2Serv:      tt.fields.kgxv2Serv,
				ninjaServ:      tt.fields.ninjaServ,
				posServ:        tt.fields.posServ,
				sicepatServ:    tt.fields.sicepatServ,
			}
			os.RunCronJob(tt.args.argument)
		}
	}
}
