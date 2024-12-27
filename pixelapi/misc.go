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

// SiaPrice is the price of one siacoin
type SiaPrice struct {
	Price float64 `json:"price"`
}

// GetSiaPrice gets the price of one siacoin
func (p *PixelAPI) GetSiaPrice() (resp float64, err error) {
	var sp SiaPrice
	return sp.Price, p.jsonRequest("GET", "misc/sia_price", &sp)
}

type RateLimits struct {
	ServerOverload    bool `json:"server_overload"`
	SpeedLimit        int  `json:"speed_limit"`
	DownloadLimit     int  `json:"download_limit"`
	DownloadLimitUsed int  `json:"download_limit_used"`
	TransferLimit     int  `json:"transfer_limit"`
	TransferLimitUsed int  `json:"transfer_limit_used"`
}

func (p *PixelAPI) GetMiscRateLimits() (rl RateLimits, err error) {
	return rl, p.jsonRequest("GET", "misc/rate_limits", &rl)
}

type ClusterSpeed struct {
	ServerTX  int64 `json:"server_tx"`
	ServerRX  int64 `json:"server_rx"`
	CacheTX   int64 `json:"cache_tx"`
	CacheRX   int64 `json:"cache_rx"`
	StorageTX int64 `json:"storage_tx"`
	StorageRX int64 `json:"storage_rx"`
}

func (p *PixelAPI) GetMiscClusterSpeed() (s ClusterSpeed, err error) {
	return s, p.jsonRequest("GET", "misc/cluster_speed", &s)
}
