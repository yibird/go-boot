package v1

import "go-boot/api/v1/sys"

type ApiGroup struct {
	SysApiGroup sys.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
