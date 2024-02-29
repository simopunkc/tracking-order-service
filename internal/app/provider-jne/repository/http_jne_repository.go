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
	"strings"
	"time"
	"tracking-order-service/internal/app/domain"
	"tracking-order-service/internal/pkg/constant"
)

const getTrackingUrl = "/tracing/api/list/v1/cnote/"

type HttpJneRepository struct {
	client *http.Client
}

func NewHttpJneRepository() *HttpJneRepository {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return &HttpJneRepository{
		client: &client,
	}
}

func (hjr HttpJneRepository) Tracking(ctx context.Context, config domain.RepositoryParam) domain.ResultJneRepository {
	isProduction, err := strconv.ParseBool(os.Getenv("PRODUCTION_MODE"))
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultJneRepository{
			Hash:  "GMsYzS1rKx3M",
			Error: err,
		}
	}

	var endpoint domain.ProviderJneEndpoint
	err = json.Unmarshal([]byte(config.Order.Order.Provider.Endpoint), &endpoint)
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultJneRepository{
			Hash:  "GMrES2HyHKzB",
			Error: err,
		}
	}

	var headers domain.ProviderJneHeaders
	err = json.Unmarshal([]byte(config.Order.Order.Provider.Headers), &headers)
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultJneRepository{
			Hash:  "GMsweTWwZ6Ca",
			Error: err,
		}
	}

	var (
		endpointFinal    domain.ProviderJneEndpointFinal
		headersFinal     domain.ProviderJneHeadersFinal
		headersAuthFinal domain.ProviderJneHeadersAuthFinal
	)

	if isProduction {
		endpointFinal = endpoint.Production
		headersFinal = headers.Production
		headersAuthFinal = headers.AuthProduction
	} else {
		endpointFinal = endpoint.Staging
		headersFinal = headers.Staging
		headersAuthFinal = headers.AuthStaging
	}

	for i := 1; i <= config.Setting.MaxRetryCheck; i++ {
		url := fmt.Sprintf("%s%s%s", endpointFinal.Default, getTrackingUrl, config.Order.Order.Awb)

		body := domain.JneRequestBodyApi{
			Username: headersAuthFinal.Username,
			ApiKey:   headersAuthFinal.ApiKey,
		}

		var buf bytes.Buffer
		err = json.NewEncoder(&buf).Encode(body)
		if err != nil {
			return domain.ResultJneRepository{
				Hash:  "GM2agxesDrsp",
				Error: err,
			}
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
		if err != nil {
			return domain.ResultJneRepository{
				Hash:  "GMCDwmsqgqcP",
				Error: err,
			}
		}
		req.Header.Set("Content-Type", headersFinal.ContentType)

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
			return domain.ResultJneRepository{
				Hash:  "GMglYjUDqzyq",
				Error: err,
			}
		}
		bodyString := string(bodyBytes)

		var provider domain.JneResponseBodyApi
		err = json.Unmarshal(bodyBytes, &provider)
		if err != nil {
			log.Println(config.Order.Order.Awb, err)
			return domain.ResultJneRepository{
				Hash:  "GMTEghghucnR",
				Error: err,
			}
		}

		return domain.ResultJneRepository{
			Hash:           "GM9h7e5zuGLm",
			ObjectResponse: provider,
			RawResponse:    bodyString,
		}
	}

	return domain.ResultJneRepository{
		Hash: "GM1xOxuUuHrC",
	}
}

func (hjr HttpJneRepository) GetLastStatusOrder(orderStatus string, jneApi domain.JneResponseBodyApi) domain.ResultJneStatusOrder {
	var lastHistory domain.JneResponseBodyApiHistory
	layout := "02-01-2006 15:04"
	lastHistoryDate := time.Date(1970, time.January, 1, 0, 00, 00, 000000000, time.UTC)
	for _, history := range jneApi.History {
		apiHistoryDate, err := time.Parse(layout, history.Date)
		if err != nil {
			continue
		}
		if apiHistoryDate.After(lastHistoryDate) {
			lastHistory = history
			lastHistoryDate = apiHistoryDate
		}
	}

	var ResultJneStatusOrder domain.ResultJneStatusOrder
	switch {
	case strings.Contains(lastHistory.Desc, "PICKED UP BY"):
		ResultJneStatusOrder = domain.ResultJneStatusOrder{
			Hash:   "GMoVZnlnT0Uf",
			Status: constant.SENDING,
		}
	case strings.Contains(lastHistory.Desc, "WITH DELIVERY COURIER"):
		ResultJneStatusOrder = domain.ResultJneStatusOrder{
			Hash:   "GMMLXUEAf7Dd",
			Status: constant.ON_COURIER,
		}
	case strings.Contains(lastHistory.Desc, "DELIVERED"):
		ResultJneStatusOrder = domain.ResultJneStatusOrder{
			Hash:   "GMAjNQSV5bF8",
			Status: constant.DELIVERED,
		}
	default:
		ResultJneStatusOrder = domain.ResultJneStatusOrder{
			Hash:   "GMoyqx079tge",
			Status: orderStatus,
		}
	}

	ResultJneStatusOrder.ModifiedAt = lastHistoryDate
	return ResultJneStatusOrder
}
