package service

import (
	"encoding/json"
	"log"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/app/order/config"
	"tracking-order-service/internal/pkg/constant"
)

//go:generate moq -out order_service_mock_test.go . DatabaseReadOrderRepository DatabaseWriteOrderRepository ProviderService
type DatabaseReadOrderRepository interface {
	OrdersModel(domain.TrackingOrderParam) ([]domain.OrderTracking, error)
	RulesModel(domain.TrackingOrderParam) (domain.Settings, error)
}

type DatabaseWriteOrderRepository interface {
	BulkUpdateOrder(orders []domain.Order) (int64, error)
	BulkUpdateOrderTracking(orders []domain.OrderTracking) (int64, error)
}

type ProviderService interface {
	RunTracking(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService
}

type OrderService struct {
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

func NewOrderService(orderRepoRead DatabaseReadOrderRepository, orderRepoWrite DatabaseWriteOrderRepository, anterajaServ ProviderService, jneServ ProviderService, jntServ ProviderService, kgxv2Serv ProviderService, ninjaServ ProviderService, posServ ProviderService, sicepatServ ProviderService) *OrderService {
	return &OrderService{orderRepoRead, orderRepoWrite, anterajaServ, jneServ, jntServ, kgxv2Serv, ninjaServ, posServ, sicepatServ}
}

func (os OrderService) workerPool(jobs <-chan domain.OrderTracking, consumer chan<- domain.ResultProviderService, rules *domain.SettingsValueCheckOrder) {
	for job := range jobs {
		consumer <- os.runTask(job, rules)
	}
}

func (os OrderService) runTask(order domain.OrderTracking, rules *domain.SettingsValueCheckOrder) domain.ResultProviderService {
	var providerService domain.ResultProviderService
	switch order.Order.Provider.Slug {
	case constant.ANTERAJA:
		providerService = os.anterajaServ.RunTracking(order, rules)
	case constant.JNT:
		providerService = os.jntServ.RunTracking(order, rules)
	case constant.JNE:
		providerService = os.jneServ.RunTracking(order, rules)
	case constant.KGX:
		providerService = os.kgxv2Serv.RunTracking(order, rules)
	case constant.NINJA:
		providerService = os.ninjaServ.RunTracking(order, rules)
	case constant.POS:
		providerService = os.posServ.RunTracking(order, rules)
	case constant.SICEPAT:
		providerService = os.sicepatServ.RunTracking(order, rules)
	}

	providerService.OrderLastModified = order.Order.ModifiedAt
	providerService.OrderTrackingLastModified = order.ModifiedAt

	return providerService
}

func (oserv OrderService) RunCronJob(argument domain.TrackingOrderParam) domain.ResultOrderService {
	orders, err := oserv.orderRepoRead.OrdersModel(argument)
	if err != nil {
		log.Println(err)
		return domain.ResultOrderService{
			Hash:         "GMtJXpTDjO7R",
			TotalUpdated: 0,
		}
	}
	if len(orders) == 0 {
		log.Println("no result orders")
		return domain.ResultOrderService{
			Hash:         "GMi7Y67vyz6B",
			TotalUpdated: 0,
		}
	}

	rules, err := oserv.orderRepoRead.RulesModel(argument)
	if err != nil {
		log.Println(err)
		return domain.ResultOrderService{
			Hash:         "GMQ5qIDozYcI",
			TotalUpdated: 0,
		}
	}

	var rulesValue domain.SettingsValueCheckOrder
	err = json.Unmarshal([]byte(rules.Value), &rulesValue)
	if err != nil {
		log.Println(err)
		return domain.ResultOrderService{
			Hash:         "GM5GTbniGiv0",
			TotalUpdated: 0,
		}
	}

	numberOfWorkers := config.MaxWorkerRunTracking
	numberOfMaximumJobs := len(orders)
	producerAnteraja := make(chan domain.OrderTracking, numberOfWorkers)
	producerJne := make(chan domain.OrderTracking, numberOfWorkers)
	producerJnt := make(chan domain.OrderTracking, numberOfWorkers)
	producerKgx := make(chan domain.OrderTracking, numberOfWorkers)
	producerNinja := make(chan domain.OrderTracking, numberOfWorkers)
	producerPos := make(chan domain.OrderTracking, numberOfWorkers)
	producerSicepat := make(chan domain.OrderTracking, numberOfWorkers)
	consumer := make(chan domain.ResultProviderService, numberOfMaximumJobs)
	listUpdatedOrder := make([]domain.Order, 0)
	listUpdatedOrderTracking := make([]domain.OrderTracking, 0)

	for w := 1; w <= numberOfWorkers; w++ {
		go oserv.workerPool(producerAnteraja, consumer, &rulesValue)
		go oserv.workerPool(producerJne, consumer, &rulesValue)
		go oserv.workerPool(producerJnt, consumer, &rulesValue)
		go oserv.workerPool(producerKgx, consumer, &rulesValue)
		go oserv.workerPool(producerNinja, consumer, &rulesValue)
		go oserv.workerPool(producerPos, consumer, &rulesValue)
		go oserv.workerPool(producerSicepat, consumer, &rulesValue)
	}

	for _, order := range orders {
		switch order.Order.Provider.Slug {
		case constant.ANTERAJA:
			producerAnteraja <- order
		case constant.JNT:
			producerJnt <- order
		case constant.JNE:
			producerJne <- order
		case constant.KGX:
			producerKgx <- order
		case constant.NINJA:
			producerNinja <- order
		case constant.POS:
			producerPos <- order
		case constant.SICEPAT:
			producerSicepat <- order
		default:
			numberOfMaximumJobs -= 1
			continue
		}
	}

	close(producerAnteraja)
	close(producerJne)
	close(producerJnt)
	close(producerKgx)
	close(producerNinja)
	close(producerPos)
	close(producerSicepat)

	for i := 0; i < numberOfMaximumJobs; i++ {
		select {
		case result := <-consumer:
			if result.NewOrder.ModifiedAt.After(result.OrderTrackingLastModified) {
				listUpdatedOrderTracking = append(listUpdatedOrderTracking, result.NewOrder)
			}
			if result.NewOrder.Order.ModifiedAt.After(result.OrderLastModified) {
				listUpdatedOrder = append(listUpdatedOrder, result.NewOrder.Order)
			}
		case <-time.After(time.Second * config.MaxSecondDeadlineRunTask):
			continue
		}
	}
	close(consumer)

	var count int64 = 0
	if len(listUpdatedOrderTracking) > 0 {
		numberUpdatedOrderTracking, err := oserv.orderRepoWrite.BulkUpdateOrderTracking(listUpdatedOrderTracking)
		if err != nil {
			log.Println(err)
			return domain.ResultOrderService{
				Hash:         "GMasJ1wQqVZJ",
				TotalUpdated: 0,
			}
		}
		count += numberUpdatedOrderTracking
	}

	if len(listUpdatedOrder) > 0 {
		numberUpdatedOrder, err := oserv.orderRepoWrite.BulkUpdateOrder(listUpdatedOrder)
		if err != nil {
			log.Println(err)
			return domain.ResultOrderService{
				Hash:         "GMmPd7S1SVRB",
				TotalUpdated: 0,
			}
		}
		count += numberUpdatedOrder
	}

	return domain.ResultOrderService{
		Hash:         "GMWg8IfAyCHH",
		TotalUpdated: count,
	}
}
