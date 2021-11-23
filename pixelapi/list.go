package pixelapi

import (
	"time"
)

// ListInfo information object from the pixeldrain API
type ListInfo struct {
	Success     bool       `json:"success"`
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	DateCreated time.Time  `json:"date_created"`
	FileCount   int        `json:"file_count"`
	Files       []ListFile `json:"files"`
	CanEdit     bool       `json:"can_edit"`
}

// ListFile information object from the pixeldrain API
type ListFile struct {
	DetailHREF  string `json:"detail_href"`
	Description string `json:"description"`
	FileInfo    `json:""`
}

// GetListID get a List from the pixeldrain API
func (p *PixelAPI) GetListID(id string) (resp ListInfo, err error) {
	return resp, p.jsonRequest("GET", "list/"+id, &resp)
}
