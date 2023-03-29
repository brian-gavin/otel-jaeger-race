# otel-jaeger-race

Reproducer for a race condition in `go.opentelemetry.io/otel/exporters/jaeger`

## Steps

Build with `-race` and run

```sh
$ go run -race main.go
```
