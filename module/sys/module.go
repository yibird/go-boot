package sys

import (
	auth "go-boot/module/sys/auth"
	menu "go-boot/module/sys/menu"
	post "go-boot/module/sys/post"
	role "go-boot/module/sys/role"
	tenant "go-boot/module/sys/tenant"
	user "go-boot/module/sys/user"
	"go.uber.org/fx"
)

var SysModule = fx.Module("sys",
	fx.Provide(
		role.NewRoleService,
		role.NewRoleApi,
		role.NewRoleRouter,
	),
	fx.Provide(
		tenant.NewTenantService,
		tenant.NewTenantApi,
		tenant.NewTenantRouter,
	),
	fx.Provide(
		post.NewPostService,
		post.NewPostApi,
		post.NewPostRouter,
	),
	fx.Provide(
		menu.NewMenuService,
		menu.NewMenuApi,
		menu.NewMenuRouter,
	),
	fx.Provide(
		user.NewUserService,
		user.NewUserApi,
		user.NewUserRouter,
	),
	fx.Provide(
		auth.NewAuthApi,
		auth.NewAuthRouter,
	),
	fx.Provide(NewSysRouter),
)
