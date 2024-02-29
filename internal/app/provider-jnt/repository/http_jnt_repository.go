package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/pkg/constant"
)

type HttpJntRepository struct {
	client *http.Client
}

func NewHttpJntRepository() *HttpJntRepository {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return &HttpJntRepository{
		client: &client,
	}
}

func (hjr HttpJntRepository) Tracking(ctx context.Context, config domain.RepositoryParam) domain.ResultJntRepository {
	isProduction, err := strconv.ParseBool(os.Getenv("PRODUCTION_MODE"))
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultJntRepository{
			Hash:  "GMLFKKsxWn9G",
			Error: err,
		}
	}

	var endpoint domain.ProviderJntEndpoint
	err = json.Unmarshal([]byte(config.Order.Order.Provider.Endpoint), &endpoint)
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultJntRepository{
			Hash:  "GMuTDsPuXzWl",
			Error: err,
		}
	}

	var headers domain.ProviderJntHeaders
	err = json.Unmarshal([]byte(config.Order.Order.Provider.Headers), &headers)
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultJntRepository{
			Hash:  "GMXprk9kpWx1",
			Error: err,
		}
	}

	var (
		endpointFinal    domain.ProviderJntEndpointFinal
		headersAuthFinal domain.ProviderJntHeadersAuthFinal
	)

	if isProduction {
		endpointFinal = endpoint.Production
		headersAuthFinal = headers.AuthProduction
	} else {
		endpointFinal = endpoint.Staging
		headersAuthFinal = headers.AuthStaging
	}

	for i := 1; i <= config.Setting.MaxRetryCheck; i++ {
		url := endpointFinal.Track

		body := domain.JntRequestBodyApi{
			Awb:         config.Order.Order.Awb,
			Eccompanyid: headersAuthFinal.Track.Username,
		}

		var buf bytes.Buffer
		err = json.NewEncoder(&buf).Encode(body)
		if err != nil {
			return domain.ResultJntRepository{
				Hash:  "GMiqP8fy0TZR",
				Error: err,
			}
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
		if err != nil {
			return domain.ResultJntRepository{
				Hash:  "GMRpHh1jxW06",
				Error: err,
			}
		}
		req.SetBasicAuth(headersAuthFinal.Track.Username, headersAuthFinal.Track.Password)
		req.Header.Set("Content-Type", "application/json")

		resp, err := hjr.client.Do(req)
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
			return domain.ResultJntRepository{
				Hash:  "GMJBrsRe0DUc",
				Error: err,
			}
		}
		bodyString := string(bodyBytes)

		var provider domain.JntResponseBodyApi
		err = json.Unmarshal(bodyBytes, &provider)
		if err != nil {
			log.Println(config.Order.Order.Awb, err)
			return domain.ResultJntRepository{
				Hash:  "GMKavaPu2Pgl",
				Error: err,
			}
		}

		return domain.ResultJntRepository{
			Hash:           "GMD6h0anilct",
			ObjectResponse: provider,
			RawResponse:    bodyString,
		}
	}

	return domain.ResultJntRepository{
		Hash: "GMbeQsY47X5k",
	}
}

func (hjr HttpJntRepository) GetLastStatusOrder(orderStatus string, jntApi domain.JntResponseBodyApi) domain.ResultJntStatusOrder {
	var lastHistory domain.JntResponseBodyApiHistory
	layout := "2006-01-02 15:04:05"
	lastHistoryDate := time.Date(1970, time.January, 1, 0, 00, 00, 000000000, time.UTC)
	for _, history := range jntApi.History {
		apiHistoryDate, err := time.Parse(layout, history.DateTime)
		if err != nil {
			continue
		}
		if apiHistoryDate.After(lastHistoryDate) {
			lastHistory = history
			lastHistoryDate = apiHistoryDate
		}
	}

	var ResultJntStatusOrder domain.ResultJntStatusOrder
	switch {
	case lastHistory.StatusCode == 101:
		ResultJntStatusOrder = domain.ResultJntStatusOrder{
			Hash:   "GM6u1XXs4qyJ",
			Status: constant.WAITING_PICKUP,
		}
	case lastHistory.StatusCode == 100:
		ResultJntStatusOrder = domain.ResultJntStatusOrder{
			Hash:   "GMnPcdwdLUpT",
			Status: constant.SENDING,
		}
	case lastHistory.Status == "Paket akan dikirim ke alamat penerima":
		ResultJntStatusOrder = domain.ResultJntStatusOrder{
			Hash:   "GMAEHe9CIk5H",
			Status: constant.ON_COURIER,
		}
	case lastHistory.StatusCode == 250:
		ResultJntStatusOrder = domain.ResultJntStatusOrder{
			Hash:   "GMqP7DlOwuJz",
			Status: constant.DELIVERED,
		}
	default:
		ResultJntStatusOrder = domain.ResultJntStatusOrder{
			Hash:   "GMZuQnxvzcIt",
			Status: orderStatus,
		}
	}

	ResultJntStatusOrder.ModifiedAt = lastHistoryDate
	return ResultJntStatusOrder
}
