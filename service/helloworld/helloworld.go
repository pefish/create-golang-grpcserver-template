package helloworld

import (
	"_template_/proto/helloworld"
	"context"
)

type HelloWorldService struct {

}

func (this *HelloWorldService) GetResult(ctx context.Context, request *helloworld.GetResultRequest) (*helloworld.GetResultReply, error) {
	//go_logger.Logger.Info(ctx.Value("method_name"), request)
	return &helloworld.GetResultReply{
		Result: request.GetText(),
	}, nil
}
