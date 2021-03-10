package pixelapi

// Recaptcha stores the reCaptcha site key
type Recaptcha struct {
	SiteKey string `json:"site_key"`
}

// GetMiscRecaptcha gets the reCaptcha site key from the pixelapi server. If
// reCaptcha is disabled the key will be empty
func (p *PixelAPI) GetMiscRecaptcha() (resp Recaptcha, err error) {
	return resp, p.jsonRequest("GET", "misc/recaptcha", &resp)
}

// GetMiscViewToken requests a viewtoken from the server. The viewtoken is valid
// for a limited amount of time and can be used to add views to a file.
// Viewtokens can only be requested from localhost
func (p *PixelAPI) GetMiscViewToken() (resp string, err error) {
	return resp, p.jsonRequest("GET", "misc/viewtoken", &resp)
}

// SiaPrice is the price of one siacoin
type SiaPrice struct {
	Price float64 `json:"price"`
}

// GetSiaPrice gets the price of one siacoin
func (p *PixelAPI) GetSiaPrice() (resp float64, err error) {
	var sp SiaPrice
	return sp.Price, p.jsonRequest("GET", "misc/sia_price", &sp)
}
