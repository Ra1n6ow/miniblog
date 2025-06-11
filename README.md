# miniblog 项目

## GRPC

### grpcurl

```shell
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
grpcurl -plaintext localhost:6666 list
```

注意：使用 grpcurl list 需要 grpc 开启反射功能

```go
// 注册反射服务
eflection.Register(grpcSrv)
```

### grpc-gateway

```shell
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.24.0
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.24.0
```

### protoc-gen-go-defaults

https://github.com/onexstack/protoc-gen-defaults
为 message 字段提供默认值

### protoc-go-inject-tag

```shell
go install github.com/favadi/protoc-go-inject-tag@latest
```

给消息字体指定 GO 结构体标签
