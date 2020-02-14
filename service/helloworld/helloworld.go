package helloworld

import (
	"_template_/proto/helloworld"
	"context"
)

type HelloWorldService struct {

}

func (this *HelloWorldService) GetResult(ctx context.Context, request *helloworld.GetResultRequest) (*helloworld.GetResultReply, error) {
	panic(`11`)
	return &helloworld.GetResultReply{
		Result: request.GetText(),
	}, nil
}
