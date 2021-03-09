package pixelapi

import (
	"net/url"

	"fornaxian.tech/pixeldrain_server/api/restapi/apitype"
)

// GetFilesystemBuckets returns a list of buckets for the user. You need to be
// authenticated
func (p *PixelAPI) GetFilesystemBuckets() (resp []apitype.Bucket, err error) {
	return resp, p.jsonRequest("GET", "filesystem", &resp)
}

// GetFilesystemPath opens a filesystem path
func (p *PixelAPI) GetFilesystemPath(path string) (resp apitype.FilesystemPath, err error) {
	return resp, p.jsonRequest("GET", "filesystem/"+url.PathEscape(path)+"?stat", &resp)
}
