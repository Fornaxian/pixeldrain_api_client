package pixelapi

import (
	"net/url"

	"fornaxian.tech/pixeldrain_server/api/restapi/apitype"
)

// AdminGetGlobals returns the global API settings
func (p *PixelAPI) AdminGetGlobals() (resp []apitype.AdminGlobal, err error) {
	return resp, p.jsonRequest("GET", "admin/globals", &resp)
}

// AdminSetGlobals sets a global API setting
func (p *PixelAPI) AdminSetGlobals(key, value string) (err error) {
	return p.form("POST", "admin/globals", url.Values{"key": {key}, "value": {value}}, nil)
}

// AdminBlockFiles blocks files from being downloaded
func (p *PixelAPI) AdminBlockFiles(text, abuseType, reporter string) (bl apitype.AdminBlockFiles, err error) {
	return bl, p.form(
		"POST", "admin/block_files",
		url.Values{"text": {text}, "type": {abuseType}, "reporter": {reporter}},
		&bl,
	)
}
