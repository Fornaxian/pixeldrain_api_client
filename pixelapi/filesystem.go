package pixelapi

import (
	"net/url"
	"time"
)

// Bucket holds a filesystem
type Bucket struct {
	Name          string            `json:"name"`
	ID            string            `json:"id"`
	DateCreated   time.Time         `json:"date_created"`
	DateModified  time.Time         `json:"date_modified"`
	ReadPassword  string            `json:"read_password"`
	WritePassword string            `json:"write_password"`
	Properties    map[string]string `json:"properties"`
	Permissions   Permissions       `json:"permissions"`
}

// FilesystemPath contains a filesystem with a bucket and all its children
// leading up to the requested node
type FilesystemPath struct {
	Bucket  Bucket           `json:"bucket"`
	Parents []FilesystemNode `json:"parents"`
	Base    FilesystemNode   `json:"base"`
}

// FilesystemNode is the return value of the GET /filesystem/ API
type FilesystemNode struct {
	Type         string    `json:"type"`
	Path         string    `json:"path"`
	Name         string    `json:"name"`
	DateCreated  time.Time `json:"date_created"`
	DateModified time.Time `json:"date_modified"`

	// File params
	FileSize int64  `json:"file_size"`
	FileType string `json:"file_type"`

	Children []FilesystemNode `json:"children"`
}

// Permissions contains the actions a user can perform on an object
type Permissions struct {
	Create bool `json:"create"`
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

// GetFilesystemBuckets returns a list of buckets for the user. You need to be
// authenticated
func (p *PixelAPI) GetFilesystemBuckets() (resp []Bucket, err error) {
	return resp, p.jsonRequest("GET", "filesystem", &resp)
}

// GetFilesystemPath opens a filesystem path
func (p *PixelAPI) GetFilesystemPath(path string) (resp FilesystemPath, err error) {
	return resp, p.jsonRequest("GET", "filesystem/"+url.PathEscape(path)+"?stat", &resp)
}
