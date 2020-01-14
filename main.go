package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
//	"os"
	"time"
	"test4/handler"
	test4 "test4/proto/test4"
	"test4/subscriber"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func main() {

	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	tracer, closer, err := cfg.New(
		"go.micro.srv.test4",
		config.Logger(jaeger.StdLogger),
	)
	fmt.Println(err)
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.test4"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	test4.RegisterTest4Handler(service.Server(), new(handler.Test4))
	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.test4", service.Server(), new(subscriber.Test4))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.test4", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
