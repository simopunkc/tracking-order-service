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
	"strings"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/pkg/constant"
)

const getTrackingUrl = "/tracking"

type HttpKgxV2Repository struct {
	client *http.Client
}

func NewHttpKgxV2Repository() *HttpKgxV2Repository {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return &HttpKgxV2Repository{
		client: &client,
	}
}

func (hkvr HttpKgxV2Repository) Tracking(ctx context.Context, config domain.RepositoryParam) domain.ResultKgxV2Repository {
	isProduction, err := strconv.ParseBool(os.Getenv("PRODUCTION_MODE"))
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultKgxV2Repository{
			Hash:  "GMkzVKl2XkUn",
			Error: err,
		}
	}

	var endpoint domain.ProviderKgxV2Endpoint
	err = json.Unmarshal([]byte(config.Order.Order.Provider.Endpoint), &endpoint)
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultKgxV2Repository{
			Hash:  "GM4Lifxn1RUv",
			Error: err,
		}
	}

	var headers domain.ProviderKgxV2Headers
	err = json.Unmarshal([]byte(config.Order.Order.Provider.Headers), &headers)
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultKgxV2Repository{
			Hash:  "GMwqjP5rQgsc",
			Error: err,
		}
	}

	var (
		endpointFinal domain.ProviderKgxV2EndpointFinal
		headersFinal  domain.ProviderKgxV2HeadersFinal
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
			return domain.ResultKgxV2Repository{
				Hash:  "GMSC5UGgitUH",
				Error: err,
			}
		}

		q := req.URL.Query()
		q.Add("connote_code", strings.Split(config.Order.Order.Awb, ".")[0])
		req.URL.RawQuery = q.Encode()

		req.Header.Set("Content-Type", headersFinal.ContentType)
		req.Header.Set("x-api-key", headersFinal.XApiKey)

		resp, err := hkvr.client.Do(req)
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
			return domain.ResultKgxV2Repository{
				Hash:  "GMCAa2VT7Pbh",
				Error: err,
			}
		}
		bodyString := string(bodyBytes)

		var provider domain.KgxV2ResponseBodyApi
		err = json.Unmarshal(bodyBytes, &provider)
		if err != nil {
			log.Println(config.Order.Order.Awb, err)
			return domain.ResultKgxV2Repository{
				Hash:  "GMRZANZ0azTm",
				Error: err,
			}
		}

		return domain.ResultKgxV2Repository{
			Hash:           "GM5SKik1Xk4k",
			ObjectResponse: provider,
			RawResponse:    bodyString,
		}
	}

	return domain.ResultKgxV2Repository{
		Hash: "GMRPd30UJCL9",
	}
}

func (hkvr HttpKgxV2Repository) GetLastStatusOrder(orderStatus string, kgxV2Api domain.KgxV2ResponseBodyApi) domain.ResultKgxV2StatusOrder {
	var lastHistory domain.KgxV2ResponseBodyApiData
	layout := time.RFC3339
	lastHistoryDate := time.Date(1970, time.January, 1, 0, 00, 00, 000000000, time.UTC)
	for _, history := range kgxV2Api.Data {
		apiHistoryDate, err := time.Parse(layout, history.Date)
		if err != nil {
			continue
		}
		if apiHistoryDate.After(lastHistoryDate) {
			lastHistory = history
			lastHistoryDate = apiHistoryDate
		}
	}

	var ResultKgxV2StatusOrder domain.ResultKgxV2StatusOrder
	switch lastHistory.Action {
	case "Create":
		ResultKgxV2StatusOrder = domain.ResultKgxV2StatusOrder{
			Hash:   "GM5XDwgbPLF0",
			Status: constant.WAITING_PICKUP,
		}
	case "OnPick":
		ResultKgxV2StatusOrder = domain.ResultKgxV2StatusOrder{
			Hash:   "GMKaxZ97G2YA",
			Status: constant.SENDING,
		}
	case "WithDeliveryCourier":
		ResultKgxV2StatusOrder = domain.ResultKgxV2StatusOrder{
			Hash:   "GMIbXAriasvh",
			Status: constant.ON_COURIER,
		}
	case "Delivered":
		ResultKgxV2StatusOrder = domain.ResultKgxV2StatusOrder{
			Hash:   "GMNjlUU0sc4d",
			Status: constant.DELIVERED,
		}
	default:
		ResultKgxV2StatusOrder = domain.ResultKgxV2StatusOrder{
			Hash:   "GMNSuyWHLBcc",
			Status: orderStatus,
		}
	}

	ResultKgxV2StatusOrder.ModifiedAt = lastHistoryDate
	return ResultKgxV2StatusOrder
}
