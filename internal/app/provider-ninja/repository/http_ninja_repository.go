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
)

const getTrackingUrl = "/1.0/orders/tracking-events/"

type HttpNinjaRepository struct {
	client *http.Client
}

func NewHttpNinjaRepository() *HttpNinjaRepository {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return &HttpNinjaRepository{
		client: &client,
	}
}

func (hnr HttpNinjaRepository) Tracking(ctx context.Context, config domain.RepositoryParam) domain.ResultNinjaRepository {
	isProduction, err := strconv.ParseBool(os.Getenv("PRODUCTION_MODE"))
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultNinjaRepository{
			Hash:  "GM6EdnGNUIqG",
			Error: err,
		}
	}

	var endpoint domain.ProviderNinjaEndpoint
	err = json.Unmarshal([]byte(config.Order.Order.Provider.Endpoint), &endpoint)
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultNinjaRepository{
			Hash:  "GMTLzEdIr9Ll",
			Error: err,
		}
	}

	var headers domain.ProviderNinjaHeaders
	err = json.Unmarshal([]byte(config.Order.Order.Provider.Headers), &headers)
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultNinjaRepository{
			Hash:  "GMXhrcSfAXdU",
			Error: err,
		}
	}

	var (
		endpointFinal domain.ProviderNinjaEndpointFinal
		headersFinal  domain.ProviderNinjaHeadersFinal
	)

	if isProduction {
		endpointFinal = endpoint.Production
		headersFinal = headers.Production
	} else {
		endpointFinal = endpoint.Staging
		headersFinal = headers.Staging
	}

	for i := 1; i <= config.Setting.MaxRetryCheck; i++ {
		url := fmt.Sprintf("%s%s%s", endpointFinal.Default, getTrackingUrl, config.Order.Order.Awb)

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return domain.ResultNinjaRepository{
				Hash:  "GM2WthpXi7Ji",
				Error: err,
			}
		}

		req.Header.Set("Content-Type", headersFinal.ContentType)
		req.Header.Set("Authorization", headersFinal.Authorization)

		resp, err := hnr.client.Do(req)
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
			return domain.ResultNinjaRepository{
				Hash:  "GMvod8k4qaY3",
				Error: err,
			}
		}
		bodyString := string(bodyBytes)

		var provider domain.NinjaResponseBodyApi
		err = json.Unmarshal(bodyBytes, &provider)
		if err != nil {
			log.Println(config.Order.Order.Awb, err)
			return domain.ResultNinjaRepository{
				Hash:  "GMfl3SL9ruuM",
				Error: err,
			}
		}

		return domain.ResultNinjaRepository{
			Hash:           "GMWOMhKGHslA",
			ObjectResponse: provider,
			RawResponse:    bodyString,
		}
	}

	return domain.ResultNinjaRepository{
		Hash: "GM5mwb6T8Lfp",
	}
}

func (hnr HttpNinjaRepository) GetLastStatusOrder(orderTracking []domain.NinjaResponseBodyApi, ninjaApi domain.NinjaResponseBodyApi) domain.ListNinjaResponseBodyApi {
	layout := "2006-01-02T15:04:05-0700"
	apiHistoryDate, err := time.Parse(layout, ninjaApi.Timestamp)
	if err != nil {
		return domain.ListNinjaResponseBodyApi{
			Hash: "GMTlv3z6P62d",
		}
	}

	totalNinjaObject := len(orderTracking)
	lastIndex := totalNinjaObject - 1
	lastHistory := orderTracking[lastIndex]
	lastHistoryDate, err := time.Parse(layout, lastHistory.Timestamp)
	if err != nil {
		return domain.ListNinjaResponseBodyApi{
			Hash: "GMbv5exIG2CB",
		}
	}

	if apiHistoryDate.After(lastHistoryDate) {
		return domain.ListNinjaResponseBodyApi{
			Hash:       "GM6d9n3CAbDA",
			History:    append(orderTracking, ninjaApi),
			LastStatus: ninjaApi.Status,
		}
	}

	return domain.ListNinjaResponseBodyApi{
		Hash: "GMQWu5ta4nHt",
	}
}
