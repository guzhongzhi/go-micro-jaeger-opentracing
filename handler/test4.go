package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	test4 "test4/proto/test4"
)

type Test4 struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Test4) Call(ctx context.Context, req *test4.Request, rsp *test4.Response) error {
	log.Log("Received Test4.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Test4) Stream(ctx context.Context, req *test4.StreamingRequest, stream test4.Test4_StreamStream) error {
	log.Logf("Received Test4.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&test4.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Test4) PingPong(ctx context.Context, stream test4.Test4_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&test4.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
