package domain

import "time"

type ProviderJntHeadersFinal struct {
	ContentType string `json:"Content-Type"`
}

type ProviderJntHeadersAuthFinal struct {
	Order        ProviderJntHeadersAuthFinalOrder        `json:"order"`
	Track        ProviderJntHeadersAuthFinalTrack        `json:"track"`
	Tarif        ProviderJntHeadersAuthFinalTarif        `json:"tarif"`
	Cancellation ProviderJntHeadersAuthFinalCancellation `json:"cancellation"`
}

type ProviderJntHeadersAuthFinalOrder struct {
	Key      string `json:"key"`
	ApiKey   string `json:"api-key"`
	Username string `json:"username"`
}

type ProviderJntHeadersAuthFinalTrack struct {
	Password string `json:"Password"`
	Username string `json:"Username"`
}

type ProviderJntHeadersAuthFinalTarif struct {
	Key      string `json:"key"`
	Username string `json:"Username"`
}

type ProviderJntHeadersAuthFinalCancellation struct {
	Key      string `json:"key"`
	ApiKey   string `json:"api-key"`
	Username string `json:"username"`
}

type ProviderJntEndpointFinal struct {
	Order        string `json:"order"`
	Track        string `json:"track"`
	Default      string `json:"default"`
	Cancellation string `json:"cancellation"`
	TariffCheck  string `json:"tariff-check"`
}

type ProviderJntHeaders struct {
	Staging        ProviderJntHeadersFinal     `json:"staging"`
	Production     ProviderJntHeadersFinal     `json:"production"`
	AuthStaging    ProviderJntHeadersAuthFinal `json:"auth-staging"`
	AuthProduction ProviderJntHeadersAuthFinal `json:"auth-production"`
}

type ProviderJntEndpoint struct {
	Staging    ProviderJntEndpointFinal `json:"staging"`
	Production ProviderJntEndpointFinal `json:"production"`
}

type JntResponseBodyApiDetailCost struct {
	ShippingCost  int `json:"shipping_cost"`
	AddCost       int `json:"add_cost"`
	InsuranceCost int `json:"insurance_cost"`
	Cod           int `json:"cod"`
	ReturnCost    int `json:"return_cost"`
}

type JntResponseBodyApiDetailSender struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`
	Zipcode string `json:"-"`
	City    string `json:"city"`
	Geoloc  string `json:"-"`
}

type JntResponseBodyApiDetailReceiver struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`
	Zipcode string `json:"zipcode"`
	City    string `json:"city"`
	Geoloc  string `json:"-"`
}

type JntResponseBodyApiDetailDriver struct {
	Id    string `json:"-"`
	Name  string `json:"name"`
	Phone string `json:"-"`
	Photo string `json:"-"`
}

type JntResponseBodyApiDetailDelivDriver struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Photo string `json:"-"`
}

type JntResponseBodyApiDetail struct {
	ShippedDate  string                              `json:"shipped_date"`
	ServicesCode string                              `json:"services_code"`
	ServicesType string                              `json:"-"`
	ActualAmount int                                 `json:"actual_amount"`
	Weight       int                                 `json:"weight"`
	Qty          int                                 `json:"qty"`
	Itemname     string                              `json:"itemname"`
	DetailCost   JntResponseBodyApiDetailCost        `json:"detail_cost"`
	Sender       JntResponseBodyApiDetailSender      `json:"sender"`
	Receiver     JntResponseBodyApiDetailReceiver    `json:"receiver"`
	Driver       JntResponseBodyApiDetailDriver      `json:"driver"`
	DelivDriver  JntResponseBodyApiDetailDelivDriver `json:"deliv_driver"`
}

type JntResponseBodyApiHistory struct {
	DateTime         string `json:"date_time"`
	CityName         string `json:"city_name"`
	Status           string `json:"status"`
	StatusCode       int    `json:"status_code"`
	StoreName        string `json:"-"`
	NextSiteName     string `json:"nextSiteName"`
	Note             string `json:"-"`
	Receiver         string `json:"-"`
	DistributionFlag string `json:"-"`
	DriverName       string `json:"-"`
	DriverPhone      string `json:"-"`
	Presenter        string `json:"-"`
	AgentName        string `json:"-"`
	Presentername    string `json:"-"`
}

type JntResponseBodyApi struct {
	Awb     string                      `json:"awb"`
	Orderid string                      `json:"orderid"`
	Detail  JntResponseBodyApiDetail    `json:"detail"`
	History []JntResponseBodyApiHistory `json:"history,omitempty"`
}

type JntRequestBodyApi struct {
	Awb         string `json:"awb"`
	Eccompanyid string `json:"eccompanyid"`
}

type ResultJntRepository struct {
	Hash           string
	ObjectResponse JntResponseBodyApi
	RawResponse    string
	Error          error
}

type ResultJntStatusOrder struct {
	Hash       string
	Status     string
	ModifiedAt time.Time
}
