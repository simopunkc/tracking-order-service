package domain

import "time"

type ProviderSicepatHeadersFinal struct {
	ApiKey string `json:"api-key"`
}

type ProviderSicepatHeadersAuthFinal struct {
	ApiKeyPartner string `json:"api-key-partner"`
}

type ProviderSicepatEndpointFinal struct {
	Default string `json:"default"`
	Partner string `json:"partner"`
}

type ProviderSicepatHeaders struct {
	Staging        ProviderSicepatHeadersFinal     `json:"staging"`
	Production     ProviderSicepatHeadersFinal     `json:"production"`
	AuthStaging    ProviderSicepatHeadersAuthFinal `json:"auth-staging"`
	AuthProduction ProviderSicepatHeadersAuthFinal `json:"auth-production"`
}

type ProviderSicepatEndpoint struct {
	Staging    ProviderSicepatEndpointFinal `json:"staging"`
	Production ProviderSicepatEndpointFinal `json:"production"`
}

type SicepatResponseBodyApiSicepatStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type SicepatResponseBodyApiSicepatResultTrackHistory struct {
	DateTime string `json:"date_time"`
	Status   string `json:"status"`
	City     string `json:"city"`
}

type SicepatResponseBodyApiSicepatResultLastStatus struct {
	DateTime string `json:"date_time"`
	Status   string `json:"status"`
	City     string `json:"city"`
}

type SicepatResponseBodyApiSicepatResult struct {
	WaybillNumber     string                                            `json:"waybill_number"`
	Kodeasal          string                                            `json:"kodeasal"`
	Kodetujuan        string                                            `json:"kodetujuan"`
	Service           string                                            `json:"service"`
	Weight            int                                               `json:"weight"`
	Partner           string                                            `json:"partner"`
	Sender            string                                            `json:"sender"`
	SenderAddress     string                                            `json:"sender_address"`
	ReceiverAddress   string                                            `json:"receiver_address"`
	ReceiverName      string                                            `json:"receiver_name"`
	Realprice         int                                               `json:"realprice"`
	Totalprice        int                                               `json:"totalprice"`
	PODReceiver       string                                            `json:"-"`
	PODReceiverTime   string                                            `json:"-"`
	SendDate          string                                            `json:"send_date"`
	TrackHistory      []SicepatResponseBodyApiSicepatResultTrackHistory `json:"track_history,omitempty"`
	LastStatus        SicepatResponseBodyApiSicepatResultLastStatus     `json:"last_status,omitempty"`
	Perwakilan        string                                            `json:"perwakilan"`
	PopSigesitImgPath string                                            `json:"-"`
	PodSigesitImgPath string                                            `json:"-"`
	PodSignImgPath    string                                            `json:"-"`
	PodImgPath        string                                            `json:"-"`
	ManifestedImgPath string                                            `json:"-"`
}

type SicepatResponseBodyApiSicepat struct {
	Status SicepatResponseBodyApiSicepatStatus `json:"status"`
	Result SicepatResponseBodyApiSicepatResult `json:"result"`
}

type SicepatResponseBodyApi struct {
	Sicepat SicepatResponseBodyApiSicepat `json:"sicepat"`
}

type ResultSicepatRepository struct {
	Hash           string
	ObjectResponse SicepatResponseBodyApi
	RawResponse    string
	Error          error
}

type ResultSicepatStatusOrder struct {
	Hash       string
	Status     string
	ModifiedAt time.Time
}
