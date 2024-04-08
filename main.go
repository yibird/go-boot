package main

import (
	"go-boot/core"
	"go-boot/core/gorm"
	"go-boot/initialize"
	"go-boot/module/sys"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(core.NewViper),
		fx.Provide(core.NewConfig),
		fx.Provide(core.NewLogger),
		fx.Provide(gorm.NewGorm),
		fx.Provide(core.NewRedisClient),
		fx.Provide(core.NewCore),
		fx.Provide(core.NewValidator),
		fx.Provide(initialize.NewEngine),

		sys.SysModule,

		fx.Invoke(initialize.RegisterRoutes),
		fx.Invoke(initialize.StartServer),
	).Run()
}
