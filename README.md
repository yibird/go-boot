一个基于 Go1.20 、Gin 、Gorm 、Jwt 、Redis 、Viper、Zap 实现的脚手架。

## 技术栈

| 依赖名称    | 描述                                                   | 版本    |
| ----------- | ------------------------------------------------------ | ------- |
| gin         | web 框架                                               | v1.9.1  |
| casbin      | 权限认证库                                             | v2.77.2 |
| gorm        | orm 框架                                               | v1.25.5 |
| viper       | 用于读取配置文件的工具库                               | v1.17.0 |
| zap         | 高性能日志框架                                         |         |
| lumberjack  | 日志分割工具                                           | v2.2.1  |
| swag        | go swagger 集成库                                      | v1.16.2 |
| gin-swagger | gin swagger 集成库                                     | v1.6.0  |
| bigcache    | 0GC 的 go 本地缓存库                                   | v3.1.0  |
| go-redis    | go 语言的 redis 客户端                                 | v9.3.0  |
| cast        | 类型转换库                                             | v1.5.1  |
| afero       | Go 的文件系统抽象系统                                  | v1.9.5  |
| sonic       | 一个极快的 JSON 序列化和反序列化库                     | v1.10.2 |
| jwt         | go 实现的 JSON Web Token                               | v5.1.0  |
| wire        | Go 编译时依赖注入库                                    | v0.5.0  |
| copier      | 模型转换库,用于 DO、DTO、VO 领域模型的转换             | v0.4.0  |
| now         | golang 的时间工具库                                    | v1.1.5  |
| pflag       | 直接替代 Go 的 flag 库，实现 POSIX/GNU 风格的 flags    | v1.0.5  |
| excelize    | 用于操作 XLAM / XLSM / XLSX / XLTM / XLTX 文件的工具库 | v2.8.0  |
| ants        | 一个高性能的 goroutine 池                              | v2.8.2  |
| lancet      | 一个 go util 函数库,类似 Java 的 Hutool、JS 的 lodash  | v2.2.7  |
| cron        | Go 的 cron 库,用于任务调度                             | v3.0.0  |
| websocket   | Go 语言的 WebSocket 实现                               | v1.5.1  |
| fsnotify    | Go 的跨平台文件系统通知库,用于监控文件变动             | v1.7.0  |
