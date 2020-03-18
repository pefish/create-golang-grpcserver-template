
#### Start

```shell script
protoc -I proto/helloworld/ proto/helloworld/helloworld.proto --go_out=plugins=grpc:proto/helloworld


GO_CONFIG=`pwd`/config/local.yaml GO_SECRET=`pwd`/secret/local.yaml go run bin/main/main.go
```

#### 生成http路由处理代码
```shell script
protoc -Iproto/helloworld/ \
    -I$GOPATH/src \
    -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway\@v1.14.3/third_party/googleapis \
    --grpc-gateway_out=proto/helloworld --go_out=plugins=grpc:proto/helloworld \
    proto/helloworld/helloworld.proto
```

#### 测试访问http
```shell script
curl localhost:8001/v1/get-result -X POST --data '{"text":"grpchaha"}'
```

#### 生成http swagger
```shell script
protoc -Iproto/helloworld/ \
  -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway\@v1.14.3/third_party/googleapis \
  --swagger_out=proto/helloworld \
  proto/helloworld/helloworld.proto
```

### grpcurl

#### 安装 grpcurl
```shell script
go get github.com/fullstorydev/grpcurl

go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```

#### 访问
```shell script
grpcurl -plaintext 0.0.0.0:8000 list  # 列出所有服务

grpcurl -plaintext 0.0.0.0:8000 list helloworld.HelloWorld  # 列出服务的所有方法

grpcurl -plaintext 0.0.0.0:8000 describe helloworld.HelloWorld  # 查看服务

grpcurl -plaintext 0.0.0.0:8000 describe helloworld.HelloWorld.GetResult  # 查看方法

grpcurl -plaintext 0.0.0.0:8000 describe helloworld.GetResultRequest  # 查看类型

grpcurl -plaintext -d '{"text": "haha"}' 0.0.0.0:8000 helloworld.HelloWorld/GetResult  # 调用方法

grpcurl -plaintext -d @ 0.0.0.0:8000 helloworld.HelloWorld/GetResult  # 调用方法。参数从输入读取
```
