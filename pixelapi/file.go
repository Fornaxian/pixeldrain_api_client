package pixelapi

import (
	"io"
	"net/url"
	"time"
)

// FileID is returned when a file has been sucessfully uploaded
type FileIDCompat struct {
	Success bool   `json:"success"`
	ID      string `json:"id"`
}
type FileID struct {
	ID string `json:"id"`
}

// FileInfo is the public file information response
type FileInfo struct {
	Success       bool      `json:"success"`
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Size          int64     `json:"size"`
	Views         int64     `json:"views"`
	BandwidthUsed int64     `json:"bandwidth_used"`
	Downloads     int64     `json:"downloads"`
	DateUpload    time.Time `json:"date_upload"`
	DateLastView  time.Time `json:"date_last_view"`
	MimeType      string    `json:"mime_type"`
	ThumbnailHREF string    `json:"thumbnail_href"`

	Availability        string `json:"availability"`
	AvailabilityMessage string `json:"availability_message"`
	AvailabilityName    string `json:"availability_name"`

	AbuseType         string `json:"abuse_type"`
	AbuseReporterName string `json:"abuse_reporter_name"`

	CanEdit bool `json:"can_edit"`
	ShowAds bool `json:"show_ads"`
}

// FileStats contains realtime statistics for a file
type FileStats struct {
	Views     int64 `json:"views"`
	Bandwidth int64 `json:"bandwidth"`
	Downloads int64 `json:"downloads"`
}

// FileTimeSeries returns historic data for a file
type FileTimeSeries struct {
	Views     TimeSeries `json:"views"`
	Downloads TimeSeries `json:"downloads"`
	Bandwidth TimeSeries `json:"bandwidth"`
}

// TimeSeries contains data captures over a time span
type TimeSeries struct {
	Timestamps []time.Time `json:"timestamps"`
	Amounts    []int64     `json:"amounts"`
}

// GetFile makes a file download request and returns a readcloser. Don't forget
// to close it!
func (p *PixelAPI) GetFile(id string) (io.ReadCloser, error) {
	return p.getRaw("file/" + id)
}

// GetFileInfo gets the FileInfo from the pixeldrain API
func (p *PixelAPI) GetFileInfo(id string) (resp FileInfo, err error) {
	return resp, p.jsonRequest("GET", "file/"+id+"/info", &resp)
}

// PostFileView adds a view to a file
func (p *PixelAPI) PostFileView(id, viewtoken string) (err error) {
	return p.form("POST", "file/"+id+"/view", url.Values{"token": {viewtoken}}, nil)
}
