package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/pkg/constant"
)

const getTrackingUrl = "/waybill"

type HttpSicepatRepository struct {
	client *http.Client
}

func NewHttpSicepatRepository() *HttpSicepatRepository {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return &HttpSicepatRepository{
		client: &client,
	}
}

func (hsr HttpSicepatRepository) Tracking(ctx context.Context, config domain.RepositoryParam) domain.ResultSicepatRepository {
	isProduction, err := strconv.ParseBool(os.Getenv("PRODUCTION_MODE"))
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultSicepatRepository{
			Hash:  "GMSZaa65Qnqr",
			Error: err,
		}
	}

	var endpoint domain.ProviderSicepatEndpoint
	err = json.Unmarshal([]byte(config.Order.Order.Provider.Endpoint), &endpoint)
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultSicepatRepository{
			Hash:  "GMUfTKGzGCXv",
			Error: err,
		}
	}

	var headers domain.ProviderSicepatHeaders
	err = json.Unmarshal([]byte(config.Order.Order.Provider.Headers), &headers)
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultSicepatRepository{
			Hash:  "GMUhyDd6QYfU",
			Error: err,
		}
	}

	var (
		endpointFinal domain.ProviderSicepatEndpointFinal
		headersFinal  domain.ProviderSicepatHeadersFinal
	)

	if isProduction {
		endpointFinal = endpoint.Production
		headersFinal = headers.Production
	} else {
		endpointFinal = endpoint.Staging
		headersFinal = headers.Staging
	}

	for i := 1; i <= config.Setting.MaxRetryCheck; i++ {
		url := fmt.Sprintf("%s%s", endpointFinal.Default, getTrackingUrl)

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return domain.ResultSicepatRepository{
				Hash:  "GMwB81PgUDom",
				Error: err,
			}
		}

		q := req.URL.Query()
		q.Add("waybill", config.Order.Order.Awb)
		req.URL.RawQuery = q.Encode()

		req.Header.Set("api-key", headersFinal.ApiKey)

		resp, err := hsr.client.Do(req)
		if err != nil {
			log.Println(config.Order.Order.Awb, err)
			if i >= config.Setting.MaxRetryCheck {
				break
			} else {
				time.Sleep(5 * time.Second)
				continue
			}
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(config.Order.Order.Awb, err)
			return domain.ResultSicepatRepository{
				Hash:  "GMkjtSEicqIW",
				Error: err,
			}
		}
		bodyString := string(bodyBytes)

		var provider domain.SicepatResponseBodyApi
		err = json.Unmarshal(bodyBytes, &provider)
		if err != nil {
			log.Println(config.Order.Order.Awb, err)
			return domain.ResultSicepatRepository{
				Hash:  "GMMWt96Mn4IZ",
				Error: err,
			}
		}

		return domain.ResultSicepatRepository{
			Hash:           "GMNxezeMaVFv",
			ObjectResponse: provider,
			RawResponse:    bodyString,
		}
	}

	return domain.ResultSicepatRepository{
		Hash: "GMoX4Yu52tF6",
	}
}

func (hsr HttpSicepatRepository) GetLastStatusOrder(orderStatus string, sicepatApi domain.SicepatResponseBodyApi) domain.ResultSicepatStatusOrder {
	var lastHistory domain.SicepatResponseBodyApiSicepatResultTrackHistory
	layout := "2006-01-02 15:04"
	lastHistoryDate := time.Date(1970, time.January, 1, 0, 00, 00, 000000000, time.UTC)
	for _, history := range sicepatApi.Sicepat.Result.TrackHistory {
		apiHistoryDate, err := time.Parse(layout, history.DateTime)
		if err != nil {
			continue
		}
		if apiHistoryDate.After(lastHistoryDate) {
			lastHistory = history
			lastHistoryDate = apiHistoryDate
		}
	}

	var ResultSicepatStatusOrder domain.ResultSicepatStatusOrder
	switch {
	case len(sicepatApi.Sicepat.Result.TrackHistory) == 0:
		ResultSicepatStatusOrder = domain.ResultSicepatStatusOrder{
			Hash:   "GMUyvfj8t4vm",
			Status: lastHistory.Status,
		}
	case lastHistory.Status == "PICKREQ":
		ResultSicepatStatusOrder = domain.ResultSicepatStatusOrder{
			Hash:   "GMasI9LIoZUj",
			Status: constant.WAITING_PICKUP,
		}
	case lastHistory.Status == "PICK":
		ResultSicepatStatusOrder = domain.ResultSicepatStatusOrder{
			Hash:   "GMj5eEkrdWcf",
			Status: constant.SENDING,
		}
	case lastHistory.Status == "DROP":
		ResultSicepatStatusOrder = domain.ResultSicepatStatusOrder{
			Hash:   "GMEtJcToiyoL",
			Status: constant.SENDING,
		}
	case lastHistory.Status == "ANT":
		ResultSicepatStatusOrder = domain.ResultSicepatStatusOrder{
			Hash:   "GMXSUaKCFqs9",
			Status: constant.ON_COURIER,
		}
	case lastHistory.Status == "DELIVERED":
		ResultSicepatStatusOrder = domain.ResultSicepatStatusOrder{
			Hash:   "GMr2lfNp9pTH",
			Status: constant.DELIVERED,
		}
	default:
		ResultSicepatStatusOrder = domain.ResultSicepatStatusOrder{
			Hash:   "GMLNXXiXgS76",
			Status: orderStatus,
		}
	}

	ResultSicepatStatusOrder.ModifiedAt = lastHistoryDate
	return ResultSicepatStatusOrder
}
