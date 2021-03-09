package pixelapi

import (
	"fornaxian.tech/pixeldrain_server/api/restapi/apitype"
)

// GetPatreonByID returns information about a patron by the ID
func (p *PixelAPI) GetPatreonByID(id string) (resp apitype.Patron, err error) {
	return resp, p.jsonRequest("GET", "patreon/"+id, &resp)
}

// PostPatreonLink links a patreon subscription to the pixeldrain account which
// is logged into this API client
func (p *PixelAPI) PostPatreonLink(id string) (err error) {
	return p.jsonRequest("POST", "patreon/"+id+"/link_subscription", nil)
}
