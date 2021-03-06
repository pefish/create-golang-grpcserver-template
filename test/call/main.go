package main

import (
	"_template_/proto/helloworld"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	forever := make(chan string)

	// Set up a connection to the server.
	dialCtx, dialCancel := context.WithTimeout(context.Background(), 5 * time.Second)
	conn, err := grpc.DialContext(dialCtx, "localhost:8001", grpc.WithInsecure(), grpc.WithBlock())
	defer dialCancel()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewHelloWorldClient(conn)


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetResult(ctx, &helloworld.GetResultRequest{
		Text: "hello world",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("result: %v", r.GetResult())

	go func() {
		for {
			time.Sleep(2 * time.Second)
		}
	}()

	<- forever
}
