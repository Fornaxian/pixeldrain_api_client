package pixelapi

import (
	"net/url"

	"fornaxian.tech/pixeldrain_server/api/restapi/apitype"
)

// GetSubscriptionID returns the subscription object identified by the given ID
func (p *PixelAPI) GetSubscriptionID(id string) (resp apitype.Subscription, err error) {
	return resp, p.jsonRequest("GET", p.apiEndpoint+"/subscription/"+url.PathEscape(id), &resp)
}

// PostSubscriptionLink links a subscription to the logged in user account. Use
// Login() before calling this function to select the account to use. This
// action cannot be undone.
func (p *PixelAPI) PostSubscriptionLink(id string) (err error) {
	return p.jsonRequest("POST", p.apiEndpoint+"/subscription/"+url.PathEscape(id)+"/link", nil)
}
