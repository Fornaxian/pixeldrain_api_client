package pixelapi

import (
	"io"
	"net/url"

	"fornaxian.tech/pixeldrain_server/api/restapi/apitype"
)

// GetFile makes a file download request and returns a readcloser. Don't forget
// to close it!
func (p *PixelAPI) GetFile(id string) (io.ReadCloser, error) {
	return p.getRaw("file/" + id)
}

// GetFileInfo gets the FileInfo from the pixeldrain API
func (p *PixelAPI) GetFileInfo(id string) (resp apitype.FileInfo, err error) {
	return resp, p.jsonRequest("GET", "file/"+id+"/info", &resp)
}

// PostFileView adds a view to a file
func (p *PixelAPI) PostFileView(id, viewtoken string) (err error) {
	return p.form("POST", "file/"+id+"/view", url.Values{"token": {viewtoken}}, nil)
}
