#!/usr/bin/env bash

set -e

run_part1() {
  set -x
  cd ./span/
  rm -rf go.mod go.sum
  go mod init github.com/angelokurtis/go-otel/span
  go get go.opentelemetry.io/otel@latest
  go get go.opentelemetry.io/otel/trace@latest
  go mod tidy
}

run_part2() {
  set -x
  cd ./starter/
  rm -rf go.mod go.sum
  go mod init github.com/angelokurtis/go-otel/starter
  go get github.com/caarlos0/env/v11@latest
  go get github.com/go-logr/logr@latest
  go get github.com/go-logr/stdr@latest
  go get github.com/gotidy/ptr@latest
  go get github.com/prometheus/client_golang@latest
  go get github.com/stretchr/testify@latest
  go get go.opentelemetry.io/contrib/propagators/aws@latest
  go get go.opentelemetry.io/contrib/propagators/b3@latest
  go get go.opentelemetry.io/contrib/propagators/jaeger@latest
  go get go.opentelemetry.io/contrib/propagators/ot@latest
  go get go.opentelemetry.io/otel@latest
  go get go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc@latest
  go get go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp@latest
  go get go.opentelemetry.io/otel/exporters/otlp/otlptrace@latest
  go get go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc@latest
  go get go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp@latest
  go get go.opentelemetry.io/otel/exporters/prometheus@latest
  go get go.opentelemetry.io/otel/exporters/stdout/stdoutmetric@latest
  go get go.opentelemetry.io/otel/sdk@latest
  go get go.opentelemetry.io/otel/sdk/metric@latest
  GOTOOLCHAIN=local go mod tidy
}

run_part1 &
run_part2 &

wait
