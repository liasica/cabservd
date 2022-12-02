# 时光驹电柜控制系统

### 第三方库
- [DTM是一款开源的分布式事务管理器](https://www.dtm.pub/guide/start.html)
- [kratos](https://go-kratos.dev/docs)
- [wire](https://zhuanlan.zhihu.com/p/399101012)
- [dig](https://darjun.github.io/2020/02/22/godailylib/dig/)
- [Loki](https://github.com/grafana/loki)
- [grpc-go-middleware](https://www.cnblogs.com/FireworksEasyCool/p/12750339.html)
- [GoKit CLI](https://github.com/GrantZheng/kit/blob/master/README_zh.md)
- [Go 教程：使用 GO-KIT 来构建微服务（上）](https://learnku.com/go/t/36923)

### 日志内容
- 时间
- 供应商
- 事件类型 (遥测 / 遥控)
- 事件描述

### 接口定义
- 管理端
  - 增删改查电柜
  - 控制电柜
- 骑手端

### kratos
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
```

### ent
```shell
go run -mod=mod entgo.io/ent/cmd/ent init --target ./internal/ent/schema CabinetBin

```
