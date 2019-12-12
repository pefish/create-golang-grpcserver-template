
#### 本地启动

```shell
protoc -I proto/helloworld/ proto/helloworld/helloworld.proto --go_out=plugins=grpc:proto/helloworld


GO_CONFIG=`pwd`/config/local.yaml GO_SECRET=`pwd`/secret/local.yaml go run main.go
```


```shell

GO_CONFIG=`pwd`/config/local.yaml GO_SECRET=`pwd`/secret/local.yaml go run client/main.go
```

