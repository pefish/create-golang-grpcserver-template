package main

import (
	"context"
	"github.com/pefish/go-application"
	"github.com/pefish/go-config"
	"github.com/pefish/go-logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/tap"
	"net"
	helloworld "template/proto/helloworld"
	helloworld_service "template/service/helloworld"
)

type StateHandler struct {
}

func (this *StateHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	return ctx
}

func (this *StateHandler) HandleRPC(ctx context.Context, rpcStats stats.RPCStats) {

}

func (this *StateHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	return ctx
}

func (this *StateHandler) HandleConn(ctx context.Context, connStats stats.ConnStats) {

}

func main() {
	go_config.Config.MustLoadYamlConfig(go_config.Configuration{
		ConfigEnvName: `GO_CONFIG`,
		SecretEnvName: `GO_SECRET`,
	})
	go_application.Application.Debug = go_config.Config.GetString(`env`) != `prod`

	go_logger.Logger.Init(`bigdata-gateway`, ``)

	// 开启grpc服务
	address := `0.0.0.0:` + go_config.Config.GetString(`port`)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		go_logger.Logger.ErrorF("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.InTapHandle(func(ctx context.Context, info *tap.Info) (i context.Context, e error) {
		return context.WithValue(ctx, "method_name", info.FullMethodName), nil
	}), grpc.StatsHandler(&StateHandler{}))
	helloworld.RegisterHelloWorldServer(s, &helloworld_service.HelloWorldService{})
	go_logger.Logger.InfoF(`grpc server started. address: %s`, address)
	if err := s.Serve(lis); err != nil {
		go_logger.Logger.ErrorF("failed to serve: %v", err)
	}
}
