package pixelapi

import (
	"net/url"
	"time"
)

// FilesystemPath contains a filesystem with a bucket and all its children
// leading up to the requested node
type FilesystemPath struct {
	Path        []FilesystemNode `json:"path"`
	BaseIndex   int              `json:"base_index"`
	Children    []FilesystemNode `json:"children"`
	Permissions Permissions      `json:"permissions"`
}

// FilesystemNode is the return value of the GET /filesystem/ API
type FilesystemNode struct {
	Type         string    `json:"type"`
	Path         string    `json:"path"`
	Name         string    `json:"name"`
	DateCreated  time.Time `json:"date_created"`
	DateModified time.Time `json:"date_modified"`

	// File params
	FileSize  int64  `json:"file_size"`
	FileType  string `json:"file_type"`
	SHA256Sum string `json:"sha256_sum"`

	// Meta params
	ID            string            `json:"id,omitempty"`
	ReadPassword  string            `json:"read_password,omitempty"`
	WritePassword string            `json:"write_password,omitempty"`
	Properties    map[string]string `json:"properties,omitempty"`
}

// Permissions contains the actions a user can perform on an object
type Permissions struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

// GetFilesystemBuckets returns a list of filesystems for the user. You need to
// be authenticated
func (p *PixelAPI) GetFilesystems() (resp []FilesystemNode, err error) {
	return resp, p.jsonRequest("GET", "filesystem", &resp)
}

// GetFilesystemPath opens a filesystem path
func (p *PixelAPI) GetFilesystemPath(path string) (resp FilesystemPath, err error) {
	return resp, p.jsonRequest("GET", "filesystem/"+url.PathEscape(path)+"?stat", &resp)
}
