package pixelapi

import "fornaxian.tech/pixeldrain_server/api/restapi/apitype"

// GetListID get a List from the pixeldrain API
func (p *PixelAPI) GetListID(id string) (resp apitype.ListInfo, err error) {
	return resp, p.jsonRequest("GET", "list/"+id, &resp)
}
