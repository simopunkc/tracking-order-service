package main

import (
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/subosito/gotenv"

	"tracking-order-service/internal/app/database"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/app/order/handler"
	orderR "tracking-order-service/internal/app/order/repository"
	orderS "tracking-order-service/internal/app/order/service"
	anterajaR "tracking-order-service/internal/app/provider-anteraja/repository"
	anterajaS "tracking-order-service/internal/app/provider-anteraja/service"
	jneR "tracking-order-service/internal/app/provider-jne/repository"
	jneS "tracking-order-service/internal/app/provider-jne/service"
	jntR "tracking-order-service/internal/app/provider-jnt/repository"
	jntS "tracking-order-service/internal/app/provider-jnt/service"
	kgxV2R "tracking-order-service/internal/app/provider-kgx-v2/repository"
	kgxV2S "tracking-order-service/internal/app/provider-kgx-v2/service"
	ninjaR "tracking-order-service/internal/app/provider-ninja/repository"
	ninjaS "tracking-order-service/internal/app/provider-ninja/service"
	posR "tracking-order-service/internal/app/provider-pos/repository"
	posS "tracking-order-service/internal/app/provider-pos/service"
	sicepatR "tracking-order-service/internal/app/provider-sicepat/repository"
	sicepatS "tracking-order-service/internal/app/provider-sicepat/service"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	err := gotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	rawOrderNumber := flag.String("on", "", "Define a order number")
	rawNotTrackableOnly := flag.String("nto", "false", "Not Trackable Only")
	rawStartDate := flag.String("sd", "", "Define a start date order filter (DD-MM-YYYY)")
	flag.Parse()

	var orderNumber string = *rawOrderNumber
	var notTrackableOnly string = *rawNotTrackableOnly
	var startDate string = *rawStartDate
	bypassCheckStatus := false
	parseNotTrackableOnly, err := strconv.ParseBool(notTrackableOnly)
	if err != nil {
		log.Fatalln(err)
	}

	if parseNotTrackableOnly {
		bypassCheckStatus = true
	}

	if startDate != "" {
		layout := "02-01-2006"
		filterStartDate, err := time.Parse(layout, startDate)
		if err != nil {
			log.Fatalln(err)
		}
		startDate = filterStartDate.Format(time.RFC3339)
	} else if parseNotTrackableOnly {
		startDate = time.Now().AddDate(0, 0, -14).Format(time.RFC3339)
	}

	dbRead := database.NewDatabaseRead()
	databaseReadOrderRepo := orderR.NewDatabaseReadOrderRepository(dbRead)
	dbWrite := database.NewDatabaseWrite()
	databaseWriteOrderRepo := orderR.NewDatabaseWriteOrderRepository(dbWrite)
	anterajaRepo := anterajaR.NewHttpAnterajaRepository()
	anterajaService := anterajaS.NewAnterajaService(anterajaRepo)
	jneRepo := jneR.NewHttpJneRepository()
	jneService := jneS.NewJneService(jneRepo)
	jntRepo := jntR.NewHttpJntRepository()
	jntService := jntS.NewJntService(jntRepo)
	kgxV2Repo := kgxV2R.NewHttpKgxV2Repository()
	kgxV2Service := kgxV2S.NewKgxV2Service(kgxV2Repo)
	ninjaRepo := ninjaR.NewHttpNinjaRepository()
	ninjaService := ninjaS.NewNinjaService(ninjaRepo)
	posRepo := posR.NewHttpPosRepository()
	posService := posS.NewPosService(posRepo)
	sicepatRepo := sicepatR.NewHttpSicepatRepository()
	sicepatService := sicepatS.NewSicepatService(sicepatRepo)
	orderService := orderS.NewOrderService(databaseReadOrderRepo, databaseWriteOrderRepo, anterajaService, jneService, jntService, kgxV2Service, ninjaService, posService, sicepatService)
	orderHandler := handler.NewOrderHandler(orderService)

	argument := domain.TrackingOrderParam{
		OrderNumber:       orderNumber,
		NotTrackableOnly:  parseNotTrackableOnly,
		BypassCheckStatus: bypassCheckStatus,
		StartDate:         startDate,
	}
	orderHandler.PrintCronJob(argument)
}
