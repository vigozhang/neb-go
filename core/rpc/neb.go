package rpc

import (
	"github.com/vigozhang/neb-go/utils/httprequest"
)

type Neb struct {
	HttpRequest *httprequest.HttpRequest
	Api         *Api
	Admin       *Admin
}

func NewNeb(request *httprequest.HttpRequest) *Neb {
	neb := &Neb{HttpRequest: request}
	neb.Api = NewApi(neb)
	neb.Admin = NewAdmin(neb)
	return neb
}

func (neb *Neb) SetRequest(request *httprequest.HttpRequest) {
	neb.HttpRequest = request
	neb.Api.SetRequest(request)
	neb.Admin.SetRequest(request)
}
