package pixelapi

import "time"

// Patron is a backer on pixeldrain's patreon campaign
type Patron struct {
	PatreonUserID           string           `json:"patreon_user_id"`
	FullName                string           `json:"full_name"`
	LastChargeDate          time.Time        `json:"last_charge_date"`
	LastChargeStatus        string           `json:"last_charge_status"`
	LifetimeSupportCents    int              `json:"lifetime_support_cents"`
	PatronStatus            string           `json:"patron_status"`
	PledgeAmountCents       int              `json:"pledge_amount_cents"`
	PledgeRelationshipStart time.Time        `json:"pledge_relationship_start"`
	UserEmail               string           `json:"user_email"`
	Subscription            SubscriptionType `json:"subscription"`
}

// GetPatreonByID returns information about a patron by the ID
func (p *PixelAPI) GetPatreonByID(id string) (resp Patron, err error) {
	return resp, p.jsonRequest("GET", "patreon/"+id, &resp)
}

// PostPatreonLink links a patreon subscription to the pixeldrain account which
// is logged into this API client
func (p *PixelAPI) PostPatreonLink(id string) (err error) {
	return p.jsonRequest("POST", "patreon/"+id+"/link_subscription", nil)
}
