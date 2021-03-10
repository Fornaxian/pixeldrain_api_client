package pixelapi

import (
	"net/url"
	"time"
)

// AdminGlobal is a global setting in pixeldrain's back-end
type AdminGlobal struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// AdminBlockFiles is an array of files which were blocked
type AdminBlockFiles struct {
	FilesBlocked []string `json:"files_blocked"`
}

// AdminAbuseReporter is an e-mail address which is allowed to send abuse
// reports to abuse@pixeldrain.com
type AdminAbuseReporter struct {
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Created      time.Time `json:"created"`
	FilesBlocked int       `json:"files_blocked"`
	LastUsed     time.Time `json:"last_used"`
}

// AdminGetGlobals returns the global API settings
func (p *PixelAPI) AdminGetGlobals() (resp []AdminGlobal, err error) {
	return resp, p.jsonRequest("GET", "admin/globals", &resp)
}

// AdminSetGlobals sets a global API setting
func (p *PixelAPI) AdminSetGlobals(key, value string) (err error) {
	return p.form("POST", "admin/globals", url.Values{"key": {key}, "value": {value}}, nil)
}

// AdminBlockFiles blocks files from being downloaded
func (p *PixelAPI) AdminBlockFiles(text, abuseType, reporter string) (bl AdminBlockFiles, err error) {
	return bl, p.form(
		"POST", "admin/block_files",
		url.Values{"text": {text}, "type": {abuseType}, "reporter": {reporter}},
		&bl,
	)
}
