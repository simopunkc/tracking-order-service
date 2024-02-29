package domain

import "time"

type ProviderAnterajaHeadersFinal struct {
	ContentType     string `json:"Content-Type"`
	AccessKeyId     string `json:"access-key-id"`
	SecretAccessKey string `json:"secret-access-key"`
}

type ProviderAnterajaEndpointFinal struct {
	Default string `json:"default"`
}

type ProviderAnterajaHeaders struct {
	Staging    ProviderAnterajaHeadersFinal `json:"staging"`
	Production ProviderAnterajaHeadersFinal `json:"production"`
}

type ProviderAnterajaEndpoint struct {
	Staging    ProviderAnterajaEndpointFinal `json:"staging"`
	Production ProviderAnterajaEndpointFinal `json:"production"`
}

type AnterajaResponseBodyApiContentHistoryMessage struct {
	Id string `json:"id"`
}

type AnterajaResponseBodyApiContentHistory struct {
	ImageUrl     string                                       `json:"-"`
	HubName      string                                       `json:"-"`
	Message      AnterajaResponseBodyApiContentHistoryMessage `json:"message"`
	Params       string                                       `json:"-"`
	TrackingCode int                                          `json:"tracking_code"`
	Timestamp    string                                       `json:"timestamp"`
}

type AnterajaResponseBodyApiContentOrderShipper struct {
	AddressNotes string `json:"-"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	Postcode     string `json:"postcode"`
}

type AnterajaResponseBodyApiContentOrderReceiver struct {
	AddressNotes string `json:"-"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	Postcode     string `json:"postcode"`
}

type AnterajaResponseBodyApiContentOrderActualShipper struct {
	ProofImages    []string `json:"-"`
	Name           string   `json:"-"`
	ProofImagesUrl []string `json:"-"`
	Relationship   string   `json:"-"`
}

type AnterajaResponseBodyApiContentOrder struct {
	BookingId      string                                           `json:"booking_id"`
	Shipper        AnterajaResponseBodyApiContentOrderShipper       `json:"shipper"`
	Waybill        string                                           `json:"waybill"`
	Receiver       AnterajaResponseBodyApiContentOrderReceiver      `json:"receiver"`
	ServiceFee     int                                              `json:"service_fee"`
	Weight         int                                              `json:"weight"`
	ServiceCode    string                                           `json:"service_code"`
	Invoice        string                                           `json:"invoice"`
	ActualShipper  AnterajaResponseBodyApiContentOrderActualShipper `json:"actual_shipper"`
	ActualReceiver string                                           `json:"-"`
}

type AnterajaResponseBodyApiContent struct {
	WaybillNo string                                  `json:"waybill_no"`
	History   []AnterajaResponseBodyApiContentHistory `json:"history,omitempty"`
	Order     AnterajaResponseBodyApiContentOrder     `json:"order"`
}

type AnterajaResponseBodyApi struct {
	Status  int                            `json:"status"`
	Info    string                         `json:"info"`
	Content AnterajaResponseBodyApiContent `json:"content"`
}

type AnterajaRequestBodyApi struct {
	WaybillNo string `json:"waybill_no"`
}

type ResultAnterajaRepository struct {
	Hash           string
	ObjectResponse AnterajaResponseBodyApi
	RawResponse    string
	Error          error
}

type ResultAnterajaStatusOrder struct {
	Hash       string
	Status     string
	ModifiedAt time.Time
}
