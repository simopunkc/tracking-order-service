package domain

import "time"

type ProviderJneHeadersFinal struct {
	ContentType string `json:"Content-Type"`
}

type ProviderJneHeadersAuthFinal struct {
	ApiKey     string `json:"api_key"`
	Username   string `json:"username"`
	MerchantId string `json:"MERCHANT_ID"`
	OlshopCust string `json:"OLSHOP_CUST"`
}

type ProviderJneEndpointFinal struct {
	Default string `json:"default"`
}

type ProviderJneHeaders struct {
	Staging        ProviderJneHeadersFinal     `json:"staging"`
	Production     ProviderJneHeadersFinal     `json:"production"`
	AuthStaging    ProviderJneHeadersAuthFinal `json:"auth-staging"`
	AuthProduction ProviderJneHeadersAuthFinal `json:"auth-production"`
}

type ProviderJneEndpoint struct {
	Staging    ProviderJneEndpointFinal `json:"staging"`
	Production ProviderJneEndpointFinal `json:"production"`
}

type JneResponseBodyApiCnote struct {
	CnoteNo           string `json:"cnote_no"`
	ReferenceNumber   string `json:"reference_number"`
	CnoteOrigin       string `json:"cnote_origin"`
	CnoteDestination  string `json:"cnote_destination"`
	CnoteServicesCode string `json:"cnote_services_code"`
	ServiceType       string `json:"service_type"`
	CnoteCustNo       string `json:"cnote_cust_no"`
	CnoteDate         string `json:"cnote_date"`
	CnotePodReceiver  string `json:"cnote_pod_receiver"`
	CnoteReceiverName string `json:"cnote_receiver_name"`
	CityName          string `json:"city_name"`
	CnotePodDate      string `json:"cnote_pod_date"`
	PodStatus         string `json:"pod_status"`
	LastStatus        string `json:"last_status"`
	CustType          string `json:"cust_type"`
	CnoteAmount       string `json:"cnote_amount"`
	CnoteWeight       string `json:"cnote_weight"`
	PodCode           string `json:"pod_code"`
	Keterangan        string `json:"keterangan"`
	CnoteGoodsDescr   string `json:"cnote_goods_descr"`
	FreightCharge     string `json:"freight_charge"`
	ShippingCost      string `json:"shipping_cost"`
	Insuranceamount   string `json:"insuranceamount"`
	Priceperkg        string `json:"priceperkg"`
	Signature         string `json:"signature"`
	Photo             string `json:"photo"`
	Long              string `json:"long"`
	Lat               string `json:"lat"`
	EstimateDelivery  string `json:"estimate_delivery"`
}

type JneResponseBodyApiDetail struct {
	CnoteNo            string `json:"cnote_no"`
	CnoteDate          string `json:"cnote_date"`
	CnoteWeight        string `json:"cnote_weight"`
	CnoteOrigin        string `json:"cnote_origin"`
	CnoteShipperName   string `json:"cnote_shipper_name"`
	CnoteShipperAddr1  string `json:"cnote_shipper_addr1"`
	CnoteShipperAddr2  string `json:"-"`
	CnoteShipperAddr3  string `json:"-"`
	CnoteShipperCity   string `json:"cnote_shipper_city"`
	CnoteReceiverName  string `json:"cnote_receiver_name"`
	CnoteReceiverAddr1 string `json:"cnote_receiver_addr1"`
	CnoteReceiverAddr2 string `json:"-"`
	CnoteReceiverAddr3 string `json:"-"`
	CnoteReceiverCity  string `json:"cnote_receiver_city"`
}

type JneResponseBodyApiHistory struct {
	Date string `json:"date"`
	Desc string `json:"desc"`
	Code string `json:"code"`
}

type JneResponseBodyApi struct {
	Cnote   JneResponseBodyApiCnote     `json:"cnote"`
	Detail  []JneResponseBodyApiDetail  `json:"detail,omitempty"`
	History []JneResponseBodyApiHistory `json:"history,omitempty"`
}

type JneRequestBodyApi struct {
	Username string `json:"username"`
	ApiKey   string `json:"api_key"`
}

type ResultJneRepository struct {
	Hash           string
	ObjectResponse JneResponseBodyApi
	RawResponse    string
	Error          error
}

type ResultJneStatusOrder struct {
	Hash       string
	Status     string
	ModifiedAt time.Time
}
