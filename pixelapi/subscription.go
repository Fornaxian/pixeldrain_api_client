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
	ID                string `json:"id"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	DisableAdDisplay  bool   `json:"disable_ad_display"`
	DisableAdsOnFiles bool   `json:"disable_ads_on_files"`
	FileSizeLimit     int64  `json:"file_size_limit"`
	FileExpiryDays    int    `json:"file_expiry_days"`
}

// GetSubscriptionID returns the subscription object identified by the given ID
func (p *PixelAPI) GetSubscriptionID(id string) (resp Subscription, err error) {
	return resp, p.jsonRequest("GET", p.apiEndpoint+"/subscription/"+url.PathEscape(id), &resp)
}

// PostSubscriptionLink links a subscription to the logged in user account. Use
// Login() before calling this function to select the account to use. This
// action cannot be undone.
func (p *PixelAPI) PostSubscriptionLink(id string) (err error) {
	return p.jsonRequest("POST", p.apiEndpoint+"/subscription/"+url.PathEscape(id)+"/link", nil)
}
