package domain

import "time"

type Order struct {
	ID                  int64     `json:"id" gorm:"primaryKey"`
	OrderNumber         string    `json:"order_number"`
	OrderNumberOriginal string    `json:"order_number_original"`
	OrderType           string    `json:"order_type"`
	OrderDate           time.Time `json:"order_date"`
	Insurance           bool      `json:"insurance"`
	Awb                 string    `json:"awb"`
	ShipperName         string    `json:"shipper_name"`
	ShipperPhone        string    `json:"shipper_phone"`
	ShipperAddress      string    `json:"shipper_address"`
	ShipperEmail        string    `json:"shipper_email"`
	ReceiverName        string    `json:"receiver_name"`
	ReceiverPhone       string    `json:"receiver_phone"`
	ReceiverAddress     string    `json:"receiver_address"`
	ReceiverEmail       string    `json:"receiver_email"`
	Qty                 int       `json:"qty"`
	KoliWeight          int       `json:"koli_weight"`
	KoliLength          int       `json:"koli_length"`
	KoliWidth           int       `json:"koli_width"`
	KoliHeight          int       `json:"koli_height"`
	KoliVolume          int       `json:"koli_volume"`
	ServiceType         string    `json:"service_type"`
	RequestBody         string    `json:"request_body,omitempty"`
	ResponseBody        string    `json:"response_body,omitempty"`
	StatusCode          string    `json:"status_code"`
	Status              string    `json:"status"`
	CreatedAt           time.Time `json:"created_at"`
	ModifiedAt          time.Time `json:"modified_at"`
	MerchantID          int64     `json:"merchant_id" gorm:"index"`
	ProviderID          int64     `json:"provider_id" gorm:"index"`
	ReceiverAreaID      string    `json:"receiver_area_id"`
	ShipperAreaID       string    `json:"shipper_area_id"`
	RequestProviderBody string    `json:"request_provider_body,omitempty"`
	Activity            string    `json:"activity,omitempty"`
	Meta                string    `json:"meta,omitempty"`
	Provider            Provider  `json:"provider" gorm:"foreignKey:ProviderID"`
	// OrderTracking       OrderTracking `json:"order_tracking,omitempty"  gorm:"foreignKey:ID"`
}

type Merchant struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	UserID     int64     `json:"user_id" gorm:"index"`
}

type MerchantApi struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	ApiKey     string    `json:"api_key"`
	Salt       string    `json:"salt"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	MerchantID int64     `json:"merchant_id" gorm:"index"`
}

type MerchantProvider struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	MerchantID int64     `json:"merchant_id" gorm:"index"`
	ProviderID int64     `json:"provider_id" gorm:"index"`
}

type Provider struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Slug       string    `json:"slug"`
	Headers    string    `json:"headers,omitempty"`
	Endpoint   string    `json:"endpoint,omitempty"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type OrderTracking struct {
	ID         int64     `json:"id,omitempty" gorm:"primaryKey"`
	Response   string    `json:"response,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	OrderID    int64     `json:"order_id" gorm:"index"`
	Order      Order     `json:"order,omitempty"  gorm:"foreignKey:OrderID"`
}

type SettingsValueCheckOrder struct {
	MaxRetryCheck    int `json:"max_retry_check"`
	MaxDayTracking   int `json:"max_day_tracking"`
	MaxDayOrderStack int `json:"max_day_order_stack"`
}

type Settings struct {
	ID         int64     `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	Value      string    `json:"value,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type TrackingOrderParam struct {
	OrderNumber       string
	NotTrackableOnly  bool
	BypassCheckStatus bool
	StartDate         string
}

type RepositoryParam struct {
	Order   OrderTracking
	Setting *SettingsValueCheckOrder
	Direct  bool
	Raw     bool
}

type ResultOrderService struct {
	Hash         string
	TotalUpdated int64
}

type ResultProviderService struct {
	Hash                      string
	NewOrder                  OrderTracking
	OrderTrackingLastModified time.Time
	OrderLastModified         time.Time
}
