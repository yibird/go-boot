一个基于 Go1.20 、Gin 、Gorm 、Jwt 、Redis 、Viper、Zap 实现的脚手架。

## 技术栈

| 依赖名称    | 描述                                                   | 版本    |
| ----------- | ------------------------------------------------------ | ------- |
| gin         | web 框架                                               | v1.9.1  |
| casbin      | 权限认证库                                             | v2.77.2 |
| gorm        | orm 框架                                               | v1.25.5 |
| viper       | 用于读取配置文件的工具库                               | v1.17.0 |
| zap         | 高性能日志框架                                         | v1.26.0 |
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

## git 提交规范

git commit 格式:

```
# git commit 提交格式
git commit -m "type(scope) : subject"
# 例子1
git commit -m "feat: 新增Butto组件"
```

commit 主题:

- feat(feature):表示添加新功能或功能增强。
- fix:表示修复 bug 或问题。
- docs:表示只涉及文档的更改,如更新文档、添加注释等。
- style:表示对代码风格、格式进行修改,不影响代码逻辑(例如，空格、缩进、分号等的更改)。
- refactor:表示对代码进行重构,既不是修复 bug 也不是添加新功能的修改。
- perf(performance):表示性能优化的修改。
- test:表示添加、修改或删除测试相关的代码。
- chore:表示对构建过程或辅助工具和库的更改,不影响生产代码(例如，更新构建脚本、配置文件等)。
- build:表示与构建系统相关的更改,如更新依赖、版本管理工具的更改等。
- ci(continuous integration):表示与持续集成流程相关的更改。
- revert:表示撤销之前的提交。
- merge:表示合并分支的工作流类型,通常用于合并开发分支或特性分支到主分支。
- release:表示发布工作流类型,用于版本发布或管理。
- hotfix:表示热修复工作流类型,通常用于紧急修复生产环境中的问题。
- init:表示初始化工作流类型,用于项目初始化或初始提交。
- workflow:表示提交的工作流或过程类型,以便更好地理解提交的上下文和目的。
- wip:wip 是 Work In Progress 的缩写,即工作正在进行中。表示提交仍然处于开发或尚未完成的状态。
