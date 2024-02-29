package domain

type ProviderNinjaHeadersFinal struct {
	ContentType   string `json:"Content-Type"`
	Authorization string `json:"authorization"`
}

type ProviderNinjaHeadersAuthFinal struct {
	Order ProviderNinjaHeadersAuthFinalOrder `json:"order"`
}

type ProviderNinjaHeadersAuthFinalOrder struct {
	ClientId     string `json:"client_id"`
	GrantType    string `json:"grant_type"`
	ClientSecret string `json:"client_secret"`
}

type ProviderNinjaEndpointFinal struct {
	Default string `json:"default"`
}

type ProviderNinjaHeaders struct {
	Staging        ProviderNinjaHeadersFinal     `json:"staging"`
	Production     ProviderNinjaHeadersFinal     `json:"production"`
	AuthStaging    ProviderNinjaHeadersAuthFinal `json:"auth-staging"`
	AuthProduction ProviderNinjaHeadersAuthFinal `json:"auth-production"`
}

type ProviderNinjaEndpoint struct {
	Staging    ProviderNinjaEndpointFinal `json:"staging"`
	Production ProviderNinjaEndpointFinal `json:"production"`
}

type NinjaResponseBodyApiPreviousMeasurements struct {
	Width            float32 `json:"width"`
	Height           float32 `json:"height"`
	Length           float32 `json:"length"`
	Size             string  `json:"size"`
	VolumetricWeight float32 `json:"volumetric_weight"`
	MeasuredWeight   float64 `json:"measured_weight"`
}

type NinjaResponseBodyApiNewMeasurements struct {
	Width            float32 `json:"width"`
	Height           float32 `json:"height"`
	Length           float32 `json:"length"`
	Size             string  `json:"size"`
	VolumetricWeight float32 `json:"volumetric_weight"`
	MeasuredWeight   float64 `json:"measured_weight"`
}

type NinjaResponseBodyApiPod struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	IdentityNumber  string `json:"identity_number,omitempty"`
	Contact         string `json:"contact"`
	Uri             string `json:"uri,omitempty"`
	LeftInSafePlace bool   `json:"left_in_safe_place"`
}

type NinjaResponseBodyApi struct {
	ShipperID            int64                                    `json:"shipper_id"`
	Status               string                                   `json:"status"`
	ShipperRefNo         string                                   `json:"shipper_ref_no,omitempty"`
	TrackingRefNo        string                                   `json:"tracking_ref_no,omitempty"`
	ShipperOrderRefNo    string                                   `json:"shipper_order_ref_no"`
	Timestamp            string                                   `json:"timestamp"`
	ID                   string                                   `json:"id,omitempty"`
	TrackingId           string                                   `json:"tracking_id"`
	PreviousStatus       string                                   `json:"previous_status,omitempty"`
	Comments             string                                   `json:"comments,omitempty"`
	PreviousSize         string                                   `json:"previous_size,omitempty"`
	NewSize              string                                   `json:"new_size,omitempty"`
	OrderId              string                                   `json:"order_id,omitempty"`
	PreviousWeight       string                                   `json:"previous_weight,omitempty"`
	NewWeight            string                                   `json:"new_weight,omitempty"`
	PreviousMeasurements NinjaResponseBodyApiPreviousMeasurements `json:"previous_measurements,omitempty"`
	NewMeasurements      NinjaResponseBodyApiNewMeasurements      `json:"new_measurements,omitempty"`
	Pod                  NinjaResponseBodyApiPod                  `json:"pod,omitempty"`
}

type ListNinjaResponseBodyApi struct {
	Hash       string                 `json:"hash,omitempty"`
	History    []NinjaResponseBodyApi `json:"history,omitempty"`
	LastStatus string                 `json:"last_status,omitempty"`
}

type ResultNinjaRepository struct {
	Hash           string
	ObjectResponse NinjaResponseBodyApi
	RawResponse    string
	Error          error
}
