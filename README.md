
## Prepare

```shell script

go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

go get github.com/grpc-ecosystem/grpc-gateway

protoc -I proto/helloworld/ -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis proto/helloworld/helloworld.proto --go_out=proto/helloworld --go-grpc_out=require_unimplemented_servers=false:proto/helloworld

```

## 生成 http 转 grpc 处理代码
```shell script
protoc -I proto/helloworld/ -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --grpc-gateway_out=proto/helloworld proto/helloworld/helloworld.proto
```

## Start

```shell
GO_CONFIG=`pwd`/config/local.yaml GO_SECRET=`pwd`/secret/local.yaml go run cmd/main/main.go
```

## 测试访问http
```shell script
curl localhost:8001/v1/get-result -X POST --data '{"text":"grpchaha"}'
```

## 生成http swagger
```shell script
protoc -I proto/helloworld/ -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --swagger_out=proto/helloworld proto/helloworld/helloworld.proto
```

## grpcurl 调试（需要服务端开启 gRPC 反射服务）

### 安装 grpcurl
```shell script
go get github.com/fullstorydev/grpcurl

go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

### 访问
```shell script
grpcurl -plaintext 0.0.0.0:8000 list  # 列出所有服务

grpcurl -plaintext 0.0.0.0:8000 list helloworld.HelloWorld  # 列出服务的所有方法

grpcurl -plaintext 0.0.0.0:8000 describe helloworld.HelloWorld  # 查看服务

grpcurl -plaintext 0.0.0.0:8000 describe helloworld.HelloWorld.GetResult  # 查看方法

grpcurl -plaintext 0.0.0.0:8000 describe helloworld.GetResultRequest  # 查看类型

grpcurl -plaintext -d '{"text": "haha"}' 0.0.0.0:8000 helloworld.HelloWorld/GetResult  # 调用方法

grpcurl -plaintext -d @ 0.0.0.0:8000 helloworld.HelloWorld/GetResult  # 调用方法。参数从输入读取
```

## Build

```shell script
make
```

