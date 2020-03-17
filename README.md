
#### 本地启动

```shell
protoc -I proto/helloworld/ proto/helloworld/helloworld.proto --go_out=plugins=grpc:proto/helloworld


GO_CONFIG=`pwd`/config/local.yaml GO_SECRET=`pwd`/secret/local.yaml go run main.go
```


```shell

GO_CONFIG=`pwd`/config/local.yaml GO_SECRET=`pwd`/secret/local.yaml go run client/main.go
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
