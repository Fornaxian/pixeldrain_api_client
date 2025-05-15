package pixelapi

import (
	"net/url"
	"time"

	"github.com/gocql/gocql"
)

// Subscription contains information about a user's subscription. When it
// started, when it ends, and what type of subscription it is
type Subscription struct {
	ID               gocql.UUID       `json:"id"`
	Used             bool             `json:"used"`
	DurationDays     int              `json:"duration_days"`
	StartTime        time.Time        `json:"start_date"`
	WarningDate      time.Time        `json:"warning_date"`
	EndDate          time.Time        `json:"end_date"`
	SubscriptionType SubscriptionType `json:"subscription_type"`
}

// SubscriptionType contains information about a subscription type. It's not the
// active subscription itself, only the properties of the subscription. Like the
// perks and cost
type SubscriptionType struct {
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	Type                   string `json:"type"`
	FileSizeLimit          int64  `json:"file_size_limit"`
	FileExpiryDays         int64  `json:"file_expiry_days"`
	StorageSpace           int64  `json:"storage_space"`
	PricePerTBStorage      int64  `json:"price_per_tb_storage"`
	PricePerTBBandwidth    int64  `json:"price_per_tb_bandwidth"`
	MonthlyTransferCap     int64  `json:"monthly_transfer_cap"`
	FileViewerBranding     bool   `json:"file_viewer_branding"`
	FilesystemAccess       bool   `json:"filesystem_access"`
	FilesystemStorageLimit int64  `json:"filesystem_storage_limit"`
}

// GetSubscriptionID returns the subscription object identified by the given ID
func (p *PixelAPI) GetSubscriptionID(id string) (resp Subscription, err error) {
	return resp, p.jsonRequest("GET", "subscription/"+url.PathEscape(id), &resp)
}

// PostSubscriptionLink links a subscription to the logged in user account. Use
// Login() before calling this function to select the account to use. This
// action cannot be undone.
func (p *PixelAPI) PostSubscriptionLink(id string) (err error) {
	return p.jsonRequest("POST", "subscription/"+url.PathEscape(id)+"/link", nil)
}

type CouponCode struct {
	ID     string `json:"id"`
	Credit int64  `json:"credit"`
	Uses   int    `json:"uses"`
}

func (p *PixelAPI) GetCouponID(id string) (resp CouponCode, err error) {
	return resp, p.jsonRequest("GET", "coupon/"+url.PathEscape(id), &resp)
}

func (p *PixelAPI) PostCouponRedeem(id string) (err error) {
	return p.jsonRequest("POST", "coupon/"+url.PathEscape(id)+"/redeem", nil)
}

type Invoice struct {
	ID             string    `json:"id"`
	Time           time.Time `json:"time"`
	Amount         int64     `json:"amount"`
	VAT            int64     `json:"vat"`
	Country        string    `json:"country"`
	PaymentGateway string    `json:"payment_gateway"`
	PaymentMethod  string    `json:"payment_method"`
	Status         string    `json:"status"`
	ProcessingFee  int64     `json:"processing_fee"`
}

func (p *PixelAPI) GetBTCPayInvoices() (resp []Invoice, err error) {
	return resp, p.jsonRequest("GET", "btcpay/invoice", &resp)
}
