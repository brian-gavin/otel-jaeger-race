package main

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/trace"
)

func init() {
	jaegerExporter, err := jaeger.New(jaeger.WithAgentEndpoint())
	if err != nil {
		panic(err)
	}
	tp := trace.NewTracerProvider(trace.WithSyncer(jaegerExporter))
	otel.SetTracerProvider(tp)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		_, span := otel.Tracer("test").Start(context.Background(), "A")
		defer span.End()
		wg.Done()
	}()
	go func() {
		_, span := otel.Tracer("test").Start(context.Background(), "B")
		defer span.End()
		wg.Done()
	}()
	wg.Wait()
}
