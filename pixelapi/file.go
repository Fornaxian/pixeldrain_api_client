package pixelapi

import (
	"io"
	"net/url"
	"time"
)

// FileID is returned when a file has been sucessfully uploaded
type FileID struct {
	ID string `json:"id"`
}

// FileInfo is the public file information response
type FileInfo struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Size              int       `json:"size"`
	Views             int       `json:"views"`
	BandwidthUsed     int       `json:"bandwidth_used"`
	BandwidthUsedPaid int       `json:"bandwidth_used_paid"`
	Downloads         int       `json:"downloads"`
	DateUpload        time.Time `json:"date_upload"`
	DateLastView      time.Time `json:"date_last_view"`
	MimeType          string    `json:"mime_type"`
	ThumbnailHREF     string    `json:"thumbnail_href"`
	HashSHA256        string    `json:"hash_sha256"`

	// Custom deletion options
	DeleteAfterDate      time.Time `json:"delete_after_date"`
	DeleteAfterDownloads int       `json:"delete_after_downloads"`

	// Abuse report information
	Availability        string `json:"availability"`
	AvailabilityMessage string `json:"availability_message"`
	AbuseType           string `json:"abuse_type"`
	AbuseReporterName   string `json:"abuse_reporter_name"`

	// Personalization
	Branding map[string]string `json:"branding,omitempty"`

	// Based on user permissions
	CanEdit            bool `json:"can_edit"`
	CanDownload        bool `json:"can_download"`
	ShowAds            bool `json:"show_ads"`
	AllowVideoPlayer   bool `json:"allow_video_player"`
	DownloadSpeedLimit int  `json:"download_speed_limit"`
	SkipFileViewer     bool `json:"skip_file_viewer,omitempty"`
}

// FileStats contains realtime statistics for a file
type FileStats struct {
	Views         int `json:"views"`
	Downloads     int `json:"downloads"`
	Bandwidth     int `json:"bandwidth"`
	BandwidthPaid int `json:"bandwidth_paid"`
}

// FileTimeSeries returns historic data for a file
type FileTimeSeries struct {
	Views         TimeSeries `json:"views"`
	Downloads     TimeSeries `json:"downloads"`
	Bandwidth     TimeSeries `json:"bandwidth"`
	BandwidthPaid TimeSeries `json:"bandwidth_paid"`
}

// TimeSeries contains data captures over a time span
type TimeSeries struct {
	Timestamps []time.Time `json:"timestamps"`
	Amounts    []int       `json:"amounts"`
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
