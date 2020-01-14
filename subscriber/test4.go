package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	test4 "test4/proto/test4"
)

type Test4 struct{}

func (e *Test4) Handle(ctx context.Context, msg *test4.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *test4.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
