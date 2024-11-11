package env_test

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/angelokurtis/go-otel/starter/_test"
	"github.com/angelokurtis/go-otel/starter/internal/env"
	"github.com/angelokurtis/go-otel/starter/internal/metric"
	"github.com/angelokurtis/go-otel/starter/internal/trace"
)

func TestTrace_TracesExporter(t *testing.T) {
	t.Run("it should correctly identify the default trace exporters", func(t *testing.T) {
		variables, err := env.LookupVariables()
		require.NoError(t, err)

		exporters := env.ToTraceExporters(variables, new(fakeExporterProvider))
		assert.Equal(t, trace.Exporters{trace.ExporterOtlp}, exporters)
	})
	t.Run("it should correctly identify the specified trace exporters", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_TRACES_EXPORTER": "logging,zipkin",
		})
		defer envvars.Unset()

		variables, err := env.LookupVariables()
		require.NoError(t, err)

		exporters := env.ToTraceExporters(variables, new(fakeExporterProvider))
		assert.Equal(t, trace.Exporters{trace.ExporterLogging, trace.ExporterZipkin}, exporters)
	})
}

func TestTrace_ExporterOTLPTracesEndpoint(t *testing.T) {
	t.Run("it should be set to the default value http://localhost:4317", func(t *testing.T) {
		variables, err := env.LookupVariables()
		require.NoError(t, err)

		endpoint := url.URL(env.ToTraceEndpoint(variables))
		assert.Equal(t, "http://localhost:4317", endpoint.String())
	})
	t.Run("it should reflect the specified custom endpoint", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_TRACES_ENDPOINT": "https://otel.com:4317",
		})
		defer envvars.Unset()

		variables, err := env.LookupVariables()
		require.NoError(t, err)

		endpoint := url.URL(env.ToTraceEndpoint(variables))
		assert.Equal(t, "https://otel.com:4317", endpoint.String())
	})
	t.Run("it should prioritize OTEL_EXPORTER_OTLP_TRACES_ENDPOINT", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_ENDPOINT":        "https://otel.com.br:4317",
			"OTEL_EXPORTER_OTLP_TRACES_ENDPOINT": "https://otel.com:4317",
		})
		defer envvars.Unset()

		variables, err := env.LookupVariables()
		require.NoError(t, err)

		endpoint := url.URL(env.ToTraceEndpoint(variables))
		assert.Equal(t, "https://otel.com:4317", endpoint.String())
	})
	t.Run("it should use OTEL_EXPORTER_OTLP_ENDPOINT as the OTLP traces endpoint", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_ENDPOINT": "https://otel.com.br:4317",
		})
		defer envvars.Unset()

		variables, err := env.LookupVariables()
		require.NoError(t, err)

		endpoint := url.URL(env.ToTraceEndpoint(variables))
		assert.Equal(t, "https://otel.com.br:4317", endpoint.String())
	})
}

func TestTrace_ExporterOTLPTracesTimeout(t *testing.T) {
	t.Run("it should be set to the default value of 10 seconds", func(t *testing.T) {
		variables, err := env.LookupVariables()
		require.NoError(t, err)

		timeout := time.Duration(env.ToTraceTimeout(variables))

		assert.Equal(t, 10*time.Second, timeout)
	})
	t.Run("it should reflect the specified custom timeout", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_TIMEOUT": "1m",
		})
		defer envvars.Unset()

		variables, err := env.LookupVariables()
		require.NoError(t, err)

		timeout := time.Duration(env.ToTraceTimeout(variables))

		assert.Equal(t, 1*time.Minute, timeout)
	})
	t.Run("it should prioritize OTEL_EXPORTER_OTLP_TRACES_TIMEOUT", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_TIMEOUT":        "10m",
			"OTEL_EXPORTER_OTLP_TRACES_TIMEOUT": "1h",
		})
		defer envvars.Unset()

		variables, err := env.LookupVariables()
		require.NoError(t, err)

		timeout := time.Duration(env.ToTraceTimeout(variables))

		assert.Equal(t, 60*time.Minute, timeout)
	})
}

func TestTrace_ExporterOTLPTracesProtocol(t *testing.T) {
	t.Run("it should be set to the default value gRPC", func(t *testing.T) {
		variables, err := env.LookupVariables()
		require.NoError(t, err)

		protocol := env.ToTraceProtocol(variables)
		assert.Equal(t, trace.ProtocolGrpc, protocol)
	})
	t.Run("it should reflect the specified custom protocol", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_TRACES_PROTOCOL": "http/protobuf",
		})
		defer envvars.Unset()

		variables, err := env.LookupVariables()
		require.NoError(t, err)

		protocol := env.ToTraceProtocol(variables)
		assert.Equal(t, trace.ProtocolHttpProtobuf, protocol)
	})
	t.Run("it should prioritize OTEL_EXPORTER_OTLP_TRACES_PROTOCOL", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_PROTOCOL":        "http/protobuf",
			"OTEL_EXPORTER_OTLP_TRACES_PROTOCOL": "grpc",
		})
		defer envvars.Unset()

		variables, err := env.LookupVariables()
		require.NoError(t, err)

		protocol := env.ToTraceProtocol(variables)
		assert.Equal(t, trace.ProtocolGrpc, protocol)
	})
}

func TestTrace_ExporterOTLPTracesCompression(t *testing.T) {
	t.Run("it should be set to the default value Gzip", func(t *testing.T) {
		variables, err := env.LookupVariables()
		require.NoError(t, err)

		compression := env.ToTraceCompression(variables)
		assert.Equal(t, trace.CompressionGzip, compression)
	})
	t.Run("it should reflect the specified custom compression", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_TRACES_COMPRESSION": "none",
		})
		defer envvars.Unset()

		variables, err := env.LookupVariables()
		require.NoError(t, err)

		compression := env.ToTraceCompression(variables)
		assert.Equal(t, trace.CompressionNone, compression)
	})
}

type fakeExporterProvider struct{}

func (f *fakeExporterProvider) TraceExporters() (trace.Exporters, bool) {
	return nil, false
}

func (f *fakeExporterProvider) MetricExporters() (metric.Exporters, bool) {
	return nil, false
}
