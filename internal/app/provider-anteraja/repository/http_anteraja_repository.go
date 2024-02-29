package repository

import (
	"bytes"
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

const getTrackingUrl = "/tracking"

type HttpAnterajaRepository struct {
	client *http.Client
}

func NewHttpAnterajaRepository() *HttpAnterajaRepository {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return &HttpAnterajaRepository{
		client: &client,
	}
}

func (har HttpAnterajaRepository) Tracking(ctx context.Context, config domain.RepositoryParam) domain.ResultAnterajaRepository {
	isProduction, err := strconv.ParseBool(os.Getenv("PRODUCTION_MODE"))
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultAnterajaRepository{
			Hash:  "GMaeCnS0w14Q",
			Error: err,
		}
	}

	var endpoint domain.ProviderAnterajaEndpoint
	err = json.Unmarshal([]byte(config.Order.Order.Provider.Endpoint), &endpoint)
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultAnterajaRepository{
			Hash:  "GMeQREg3ZEjy",
			Error: err,
		}
	}

	var headers domain.ProviderAnterajaHeaders
	err = json.Unmarshal([]byte(config.Order.Order.Provider.Headers), &headers)
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultAnterajaRepository{
			Hash:  "GMI6lgzHhfoA",
			Error: err,
		}
	}

	var (
		endpointFinal domain.ProviderAnterajaEndpointFinal
		headersFinal  domain.ProviderAnterajaHeadersFinal
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

		body := domain.AnterajaRequestBodyApi{
			WaybillNo: config.Order.Order.Awb,
		}

		var buf bytes.Buffer
		err = json.NewEncoder(&buf).Encode(body)
		if err != nil {
			log.Println(config.Order.Order.Awb, err)
			return domain.ResultAnterajaRepository{
				Hash:  "GMgHWVQLTGbD",
				Error: err,
			}
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
		if err != nil {
			log.Println(config.Order.Order.Awb, err)
			return domain.ResultAnterajaRepository{
				Hash:  "GMLRiWrwpYvj",
				Error: err,
			}
		}
		req.Header.Set("Content-Type", headersFinal.ContentType)
		req.Header.Set("access-key-id", headersFinal.AccessKeyId)
		req.Header.Set("secret-access-key", headersFinal.SecretAccessKey)

		resp, err := har.client.Do(req)
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
			return domain.ResultAnterajaRepository{
				Hash:  "GM05M6dprxje",
				Error: err,
			}
		}
		bodyString := string(bodyBytes)

		var provider domain.AnterajaResponseBodyApi
		err = json.Unmarshal(bodyBytes, &provider)
		if err != nil {
			log.Println(config.Order.Order.Awb, err)
			return domain.ResultAnterajaRepository{
				Hash:  "GMTn3Xw7GTsz",
				Error: err,
			}
		}

		return domain.ResultAnterajaRepository{
			Hash:           "GMgwIF2wZ9YC",
			ObjectResponse: provider,
			RawResponse:    bodyString,
		}
	}

	return domain.ResultAnterajaRepository{
		Hash: "GMJiBkAjLVZ3",
	}
}

func (har HttpAnterajaRepository) GetLastStatusOrder(orderStatus string, anterajaApi domain.AnterajaResponseBodyApi) domain.ResultAnterajaStatusOrder {
	var lastHistory domain.AnterajaResponseBodyApiContentHistory
	layout := "2006-01-02T15:04:05.000+0000"
	lastHistoryDate := time.Date(1970, time.January, 1, 0, 00, 00, 000000000, time.UTC)
	for _, history := range anterajaApi.Content.History {
		apiHistoryDate, err := time.Parse(layout, history.Timestamp)
		if err != nil {
			continue
		}
		if apiHistoryDate.After(lastHistoryDate) {
			lastHistory = history
			lastHistoryDate = apiHistoryDate
		}
	}

	var ResultAnterajaStatusOrder domain.ResultAnterajaStatusOrder
	switch {
	case len(anterajaApi.Content.History) == 0:
		ResultAnterajaStatusOrder = domain.ResultAnterajaStatusOrder{
			Hash:   "GMM4IkfTP27k",
			Status: constant.WAITING_PICKUP,
		}
	case lastHistory.TrackingCode == 430:
		ResultAnterajaStatusOrder = domain.ResultAnterajaStatusOrder{
			Hash:   "GMvgCDffNYgc",
			Status: constant.CANCELED,
		}
	case lastHistory.TrackingCode == 200:
		ResultAnterajaStatusOrder = domain.ResultAnterajaStatusOrder{
			Hash:   "GMsut8f6cqNr",
			Status: constant.SENDING,
		}
	case lastHistory.TrackingCode == 240:
		ResultAnterajaStatusOrder = domain.ResultAnterajaStatusOrder{
			Hash:   "GMP8jJRRuu3M",
			Status: constant.ON_COURIER,
		}
	case lastHistory.TrackingCode == 250:
		ResultAnterajaStatusOrder = domain.ResultAnterajaStatusOrder{
			Hash:   "GMfxPZ6Jxx7n",
			Status: constant.DELIVERED,
		}
	default:
		ResultAnterajaStatusOrder = domain.ResultAnterajaStatusOrder{
			Hash:   "GMieS4o9vTCI",
			Status: orderStatus,
		}
	}

	ResultAnterajaStatusOrder.ModifiedAt = lastHistoryDate
	return ResultAnterajaStatusOrder
}
