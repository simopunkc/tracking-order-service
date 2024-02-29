package domain

import "time"

type ProviderPosHeadersFinal struct {
	Accept        string `json:"accept"`
	ContentType   string `json:"Content-Type"`
	Authorization string `json:"authorization"`
}

type ProviderPosHeadersAuthFinal struct {
	Key          string `json:"key"`
	Secret       string `json:"secret"`
	Userid       int    `json:"userid"`
	Memberid     string `json:"memberid"`
	XApiKey      string `json:"X-API-KEY"`
	XPosUser     string `json:"X-POS-USER"`
	XPosPassword string `json:"X-POS-PASSWORD"`
}

type ProviderPosEndpointFinal struct {
	Default string `json:"default"`
}

type ProviderPosEndpointFinalPath struct {
	CancelOrder   string `json:"cancel-order"`
	CreateOrder   string `json:"create-order"`
	ServiceRate   string `json:"service-rate"`
	TrackingOrder string `json:"tracking-order"`
}

type ProviderPosHeaders struct {
	Staging        ProviderPosHeadersFinal     `json:"staging"`
	Production     ProviderPosHeadersFinal     `json:"production"`
	AuthStaging    ProviderPosHeadersAuthFinal `json:"auth-staging"`
	AuthProduction ProviderPosHeadersAuthFinal `json:"auth-production"`
}

type ProviderPosEndpoint struct {
	Staging        ProviderPosEndpointFinal     `json:"staging"`
	Production     ProviderPosEndpointFinal     `json:"production"`
	StagingPath    ProviderPosEndpointFinalPath `json:"staging-path"`
	ProductionPath ProviderPosEndpointFinalPath `json:"production-path"`
}

type PosResponseBodyApiToken struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type PosResponseBodyApiResponseData struct {
	Status      string `json:"status,omitempty"`
	Barcode     string `json:"barcode"`
	EventDate   string `json:"eventDate"`
	EventName   string `json:"eventName"`
	OfficeCode  string `json:"officeCode"`
	OfficeName  string `json:"officeName,omitempty"`
	Office      string `json:"office,omitempty"`
	Description string `json:"description"`
}

type PosResponseBodyApiResponse struct {
	RTnt []PosResponseBodyApiResponseData `json:"r_tnt,omitempty"`
	Data []PosResponseBodyApiResponseData `json:"data,omitempty"`
}

type PosResponseBodyApi struct {
	RsTnt    PosResponseBodyApiResponse `json:"rs_tnt,omitempty"`
	Response PosResponseBodyApiResponse `json:"response,omitempty"`
}

type PosRequestBodyApiToken struct {
	GrantType string `json:"grant_type"`
}

type PosRequestBodyApi struct {
	Barcode string `json:"barcode"`
}

type ResultPosRepository struct {
	Hash           string
	ObjectResponse PosResponseBodyApi
	RawResponse    string
	Error          error
}

type ResultPosStatusOrder struct {
	Hash       string
	Status     string
	ModifiedAt time.Time
}
