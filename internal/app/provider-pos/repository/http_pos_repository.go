package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
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

const getTokenUrl = "/token"

type HttpPosRepository struct {
	client *http.Client
}

func NewHttpPosRepository() *HttpPosRepository {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	return &HttpPosRepository{
		client: &client,
	}
}

func (hpr HttpPosRepository) Tracking(ctx context.Context, config domain.RepositoryParam) domain.ResultPosRepository {
	isProduction, err := strconv.ParseBool(os.Getenv("PRODUCTION_MODE"))
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultPosRepository{
			Hash:  "GMlcXvYtJOoi",
			Error: err,
		}
	}

	var endpoint domain.ProviderPosEndpoint
	err = json.Unmarshal([]byte(config.Order.Order.Provider.Endpoint), &endpoint)
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultPosRepository{
			Hash:  "GMPTbsvsMroL",
			Error: err,
		}
	}

	var headers domain.ProviderPosHeaders
	err = json.Unmarshal([]byte(config.Order.Order.Provider.Headers), &headers)
	if err != nil {
		log.Println(config.Order.Order.Awb, err)
		return domain.ResultPosRepository{
			Hash:  "GMN5YTvDRyOz",
			Error: err,
		}
	}

	var (
		endpointFinal     domain.ProviderPosEndpointFinal
		endpointPathFinal domain.ProviderPosEndpointFinalPath
		headersFinal      domain.ProviderPosHeadersFinal
		headersAuthFinal  domain.ProviderPosHeadersAuthFinal
	)

	if isProduction {
		endpointFinal = endpoint.Production
		endpointPathFinal = endpoint.ProductionPath
		headersFinal = headers.Production
		headersAuthFinal = headers.AuthProduction
	} else {
		endpointFinal = endpoint.Staging
		endpointPathFinal = endpoint.StagingPath
		headersFinal = headers.Staging
		headersAuthFinal = headers.AuthStaging
	}

	var token domain.PosResponseBodyApiToken

	for i := 1; i <= config.Setting.MaxRetryCheck; i++ {
		urlToken := fmt.Sprintf("%s%s", endpointFinal.Default, getTokenUrl)

		bodyToken := domain.PosRequestBodyApiToken{
			GrantType: "client_credentials",
		}

		var bufToken bytes.Buffer
		err = json.NewEncoder(&bufToken).Encode(bodyToken)
		if err != nil {
			return domain.ResultPosRepository{
				Hash:  "GMHlJiEmzyr6",
				Error: err,
			}
		}

		reqToken, err := http.NewRequestWithContext(ctx, http.MethodPost, urlToken, &bufToken)
		if err != nil {
			return domain.ResultPosRepository{
				Hash:  "GMrF35rYCV6x",
				Error: err,
			}
		}
		reqToken.SetBasicAuth(headersAuthFinal.Key, headersAuthFinal.Secret)
		reqToken.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		respToken, err := hpr.client.Do(reqToken)
		if err != nil {
			log.Println(config.Order.Order.Awb, err)
			if i == config.Setting.MaxRetryCheck {
				break
			} else {
				// time.Sleep(5 * time.Second)
				continue
			}
		}

		err = json.NewDecoder(respToken.Body).Decode(&token)
		if err != nil {
			return domain.ResultPosRepository{
				Hash:  "GMZsW0TdklMW",
				Error: err,
			}
		}

		break
	}

	if token.AccessToken == "" {
		return domain.ResultPosRepository{
			Hash:  "GMI2rmqNsx0u",
			Error: errors.New("empty token"),
		}
	}

	for j := 1; j <= config.Setting.MaxRetryCheck; j++ {
		url := fmt.Sprintf("%s%s", endpointFinal.Default, endpointPathFinal.TrackingOrder)

		body := domain.PosRequestBodyApi{
			Barcode: config.Order.Order.Awb,
		}

		var buf bytes.Buffer
		err = json.NewEncoder(&buf).Encode(body)
		if err != nil {
			return domain.ResultPosRepository{
				Hash:  "GMKgraycyjj5",
				Error: err,
			}
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
		if err != nil {
			return domain.ResultPosRepository{
				Hash:  "GMHXXInXpmLo",
				Error: err,
			}
		}
		req.Header.Set("Accept", headersFinal.Accept)
		req.Header.Set("Content-Type", headersFinal.ContentType)
		req.Header.Set("X-POS-USER", headersAuthFinal.XPosUser)
		req.Header.Set("X-POS-PASSWORD", headersAuthFinal.XPosPassword)
		if token.TokenType == "Bearer" {
			bearer := "Bearer " + token.AccessToken
			req.Header.Set("Authorization", bearer)
		}

		resp, err := hpr.client.Do(req)
		if err != nil {
			log.Println(config.Order.Order.Awb, err)
			if j >= config.Setting.MaxRetryCheck {
				break
			} else {
				time.Sleep(5 * time.Second)
				continue
			}
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(config.Order.Order.Awb, err)
			return domain.ResultPosRepository{
				Hash:  "GMnqygTuI1Bi",
				Error: err,
			}
		}
		bodyString := string(bodyBytes)

		var provider domain.PosResponseBodyApi
		err = json.Unmarshal(bodyBytes, &provider)
		if err != nil {
			log.Println(config.Order.Order.Awb, err)
			return domain.ResultPosRepository{
				Hash:  "GMzsEqQznD1H",
				Error: err,
			}
		}

		if provider.RsTnt.RTnt != nil && provider.Response.Data == nil {
			provider.Response.Data = provider.RsTnt.RTnt
		}

		return domain.ResultPosRepository{
			Hash:           "GM3OXb9SBSXR",
			ObjectResponse: provider,
			RawResponse:    bodyString,
		}
	}

	return domain.ResultPosRepository{
		Hash: "GMCHpsvRbNJX",
	}
}

func (hpr HttpPosRepository) GetLastStatusOrder(orderStatus string, posApi domain.PosResponseBodyApi) domain.ResultPosStatusOrder {
	var lastHistory domain.PosResponseBodyApiResponseData
	layout := "2006-01-02 15:04:05"
	lastHistoryDate := time.Date(1970, time.January, 1, 0, 00, 00, 000000000, time.UTC)
	for _, history := range posApi.Response.Data {
		apiHistoryDate, err := time.Parse(layout, history.EventDate)
		if err != nil {
			continue
		}
		if apiHistoryDate.After(lastHistoryDate) {
			lastHistory = history
			lastHistoryDate = apiHistoryDate
		}
	}

	var ResultPosStatusOrder domain.ResultPosStatusOrder
	switch {
	case len(posApi.Response.Data) == 0:
		ResultPosStatusOrder = domain.ResultPosStatusOrder{
			Hash:   "GMbj4oAJqyxd",
			Status: constant.WAITING_PICKUP,
		}
	case lastHistory.EventName == "PRA COLLECTING":
		ResultPosStatusOrder = domain.ResultPosStatusOrder{
			Hash:   "GMrTUhdearpr",
			Status: constant.WAITING_PICKUP,
		}
	case lastHistory.EventName == "MANIFEST":
		ResultPosStatusOrder = domain.ResultPosStatusOrder{
			Hash:   "GMdoLGSXx6iw",
			Status: constant.SENDING,
		}
	case lastHistory.EventName == "PENERIMAAN DI LOKET":
		ResultPosStatusOrder = domain.ResultPosStatusOrder{
			Hash:   "GMzmV3DJvdr5",
			Status: constant.SENDING,
		}
	case lastHistory.EventName == "DALAM PROSES":
		ResultPosStatusOrder = domain.ResultPosStatusOrder{
			Hash:   "GMB7m9Exm77t",
			Status: constant.SENDING,
		}
	case lastHistory.EventName == "PROSES ANTAR":
		ResultPosStatusOrder = domain.ResultPosStatusOrder{
			Hash:   "GMpb34488fwl",
			Status: constant.ON_COURIER,
		}
	case lastHistory.EventName == "SELESAI ANTAR":
		ResultPosStatusOrder = domain.ResultPosStatusOrder{
			Hash:   "GM8iCg4K3Idd",
			Status: constant.DELIVERED,
		}
	default:
		ResultPosStatusOrder = domain.ResultPosStatusOrder{
			Hash:   "GMjg9FTonUC4",
			Status: orderStatus,
		}
	}

	ResultPosStatusOrder.ModifiedAt = lastHistoryDate
	return ResultPosStatusOrder
}
