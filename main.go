package main

import (
	"_template_/proto/helloworld"
	helloworld_service "_template_/service/helloworld"
	"context"
	runtime2 "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pefish/go-application"
	"github.com/pefish/go-config"
	"github.com/pefish/go-logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	go_config.Config.MustLoadYamlConfig(go_config.Configuration{
		ConfigEnvName: `GO_CONFIG`,
		SecretEnvName: `GO_SECRET`,
	})
	go_application.Application.Debug = go_config.Config.GetString(`env`) != `prod`

	go_logger.Logger.Init(`default`, ``)

	// 开启grpc服务
	address := `0.0.0.0:` + go_config.Config.GetString(`port`)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		go_logger.Logger.ErrorF("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.StreamInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
			defer func() {
				if err_ := recover(); err_ != nil {
					go_logger.Logger.Error(err_)
					err = status.Errorf(codes.Internal, "%#v", err_)
				}
			}()
			go_logger.Logger.DebugF("method: %s, param: %#v", info.FullMethod, srv)
			return handler(srv, ss)
		}),
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
			defer func() {
				if err_ := recover(); err_ != nil {
					go_logger.Logger.Error(err_)
					err = status.Errorf(codes.Internal, "%#v", err_)
				}
			}()
			go_logger.Logger.DebugF("method: %s, param: %#v", info.FullMethod, req)
			return handler(ctx, req)
		}),
	)
	helloworld.RegisterHelloWorldServer(s, &helloworld_service.HelloWorldService{})

	httpPort := go_config.Config.GetString(`httpPort`)
	if httpPort != `` {
		go func() {
			// 启动http服务
			ctx := context.Background()
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()
			mux := runtime2.NewServeMux()
			err = helloworld.RegisterHelloWorldHandlerFromEndpoint(
				ctx, mux, address,
				[]grpc.DialOption{grpc.WithInsecure()},
			)
			if err != nil {
				log.Fatal(err)
			}
			httpAddress := "0.0.0.0:" + httpPort
			go_logger.Logger.InfoF(`http server started. address: %s`, httpAddress)
			err = http.ListenAndServe(httpAddress, mux)
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	go_logger.Logger.InfoF(`grpc server started. address: %s`, address)
	if err := s.Serve(lis); err != nil {
		go_logger.Logger.ErrorF("failed to serve: %v", err)
	}
}
