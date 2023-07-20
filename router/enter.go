package router

import "go-boot/router/sys"

type RouterGroup struct {
	System sys.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
