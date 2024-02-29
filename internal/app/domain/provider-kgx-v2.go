package domain

import "time"

type ProviderKgxV2HeadersFinal struct {
	XApiKey     string `json:"x-api-key"`
	ContentType string `json:"Content-Type"`
}

type ProviderKgxV2EndpointFinal struct {
	Default string `json:"default"`
}

type ProviderKgxV2Headers struct {
	Staging    ProviderKgxV2HeadersFinal `json:"staging"`
	Production ProviderKgxV2HeadersFinal `json:"production"`
}

type ProviderKgxV2Endpoint struct {
	Staging    ProviderKgxV2EndpointFinal `json:"staging"`
	Production ProviderKgxV2EndpointFinal `json:"production"`
}

type KgxV2ResponseBodyApiData struct {
	Id           int64  `json:"id"`
	Action       string `json:"action"`
	ConnoteState string `json:"connote_state"`
	Content      string `json:"content"`
	Date         string `json:"date"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	LocationName string `json:"location_name"`
	Username     string `json:"username"`
	ConnoteCode  string `json:"connote_code"`
	IsHide       int    `json:"is_hide"`
	Coordinate   string `json:"-"`
}

type KgxV2ResponseBodyApi struct {
	Data         []KgxV2ResponseBodyApiData `json:"data,omitempty"`
	From         int                        `json:"from"`
	To           int                        `json:"to"`
	Total        int                        `json:"total"`
	PerPage      int                        `json:"per_page"`
	CurrentPage  int                        `json:"current_page"`
	LastPage     int                        `json:"last_page"`
	FirstPageUrl string                     `json:"first_page_url"`
	PrevPageUrl  string                     `json:"-"`
	NextPageUrl  string                     `json:"-"`
	LastPageUrl  string                     `json:"last_page_url"`
	Path         string                     `json:"path"`
}

type ResultKgxV2Repository struct {
	Hash           string
	ObjectResponse KgxV2ResponseBodyApi
	RawResponse    string
	Error          error
}

type ResultKgxV2StatusOrder struct {
	Hash       string
	Status     string
	ModifiedAt time.Time
}
