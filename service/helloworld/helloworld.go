package helloworld

import (
	"_template_/proto/helloworld"
	"context"
	go_logger "github.com/pefish/go-logger"
)

type HelloWorldService struct {

}

func (this *HelloWorldService) GetResult(ctx context.Context, request *helloworld.GetResultRequest) (*helloworld.GetResultReply, error) {
	//panic(`11`)
	go_logger.Logger.DebugF("debug log: %s", request.Text)
	return &helloworld.GetResultReply{
		Result: request.GetText(),
	}, nil
}
