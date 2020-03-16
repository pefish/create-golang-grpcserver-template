
#### Start

```shell
protoc -I proto/helloworld/ proto/helloworld/helloworld.proto --go_out=plugins=grpc:proto/helloworld


GO_CONFIG=`pwd`/config/local.yaml GO_SECRET=`pwd`/secret/local.yaml go run bin/main/main.go
```

