package service

import "go-boot/service/sys"

type ServiceGroup struct {
	SystemServiceGroup sys.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
