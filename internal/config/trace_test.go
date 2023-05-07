package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/angelokurtis/go-starter-otel/_test"
	"github.com/angelokurtis/go-starter-otel/internal/config"
	"github.com/angelokurtis/go-starter-otel/internal/trace"
)

func TestTrace_TracesExporter(t *testing.T) {
	t.Run("", func(t *testing.T) {
		env, err := config.NewFromEnv()
		require.NoError(t, err)

		var traceConfig trace.Config = config.NewTrace(env)
		exporter := traceConfig.TracesExporter()
		assert.Equal(t, exporter, []trace.Exporter{trace.ExporterOtlp})
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_TRACES_EXPORTER": "logging,zipkin",
		})
		defer envvars.Unset()

		env, err := config.NewFromEnv()
		require.NoError(t, err)

		var traceConfig trace.Config = config.NewTrace(env)
		exporter := traceConfig.TracesExporter()
		assert.Equal(t, exporter, []trace.Exporter{trace.ExporterLogging, trace.ExporterZipkin})
	})
}

func TestTrace_ExporterOTLPTracesEndpoint(t *testing.T) {
	t.Run("", func(t *testing.T) {
		env, err := config.NewFromEnv()
		require.NoError(t, err)

		var traceConfig trace.Config = config.NewTrace(env)
		endpoint := traceConfig.ExporterOTLPTracesEndpoint()
		assert.Equal(t, endpoint.String(), "http://localhost:4317")
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"EXPORTER_OTLP_TRACES_ENDPOINT": "https://otel.com:4317",
		})
		defer envvars.Unset()

		env, err := config.NewFromEnv()
		require.NoError(t, err)

		var traceConfig trace.Config = config.NewTrace(env)
		endpoint := traceConfig.ExporterOTLPTracesEndpoint()
		assert.Equal(t, endpoint.String(), "https://otel.com:4317")
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"EXPORTER_OTLP_ENDPOINT":        "https://otel.com.br:4317",
			"EXPORTER_OTLP_TRACES_ENDPOINT": "https://otel.com:4317",
		})
		defer envvars.Unset()

		env, err := config.NewFromEnv()
		require.NoError(t, err)

		var traceConfig trace.Config = config.NewTrace(env)
		endpoint := traceConfig.ExporterOTLPTracesEndpoint()
		assert.Equal(t, endpoint.String(), "https://otel.com:4317")
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"EXPORTER_OTLP_ENDPOINT": "https://otel.com.br:4317",
		})
		defer envvars.Unset()

		env, err := config.NewFromEnv()
		require.NoError(t, err)

		var traceConfig trace.Config = config.NewTrace(env)
		endpoint := traceConfig.ExporterOTLPTracesEndpoint()
		assert.Equal(t, endpoint.String(), "https://otel.com.br:4317")
	})
}

func TestTrace_ExporterOTLPTracesTimeout(t *testing.T) {
	t.Run("", func(t *testing.T) {
		env, err := config.NewFromEnv()
		require.NoError(t, err)

		var traceConfig trace.Config = config.NewTrace(env)
		timeout := traceConfig.ExporterOTLPTracesTimeout()
		var oneSecond int64 = 1000
		assert.Equal(t, timeout.Milliseconds(), oneSecond*10)
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_TIMEOUT": "1m",
		})
		defer envvars.Unset()

		env, err := config.NewFromEnv()
		require.NoError(t, err)

		var traceConfig trace.Config = config.NewTrace(env)
		timeout := traceConfig.ExporterOTLPTracesTimeout()
		var oneMinute int64 = 60000
		assert.Equal(t, timeout.Milliseconds(), oneMinute*1)
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_TIMEOUT":        "10m",
			"OTEL_EXPORTER_OTLP_TRACES_TIMEOUT": "1h",
		})
		defer envvars.Unset()

		env, err := config.NewFromEnv()
		require.NoError(t, err)

		var traceConfig trace.Config = config.NewTrace(env)
		timeout := traceConfig.ExporterOTLPTracesTimeout()
		var oneMinute int64 = 60000
		assert.Equal(t, timeout.Milliseconds(), oneMinute*60)
	})
}

func TestTrace_ExporterOTLPTracesProtocol(t *testing.T) {
	t.Run("", func(t *testing.T) {
		env, err := config.NewFromEnv()
		require.NoError(t, err)

		var traceConfig trace.Config = config.NewTrace(env)
		protocol := traceConfig.ExporterOTLPTracesProtocol()
		assert.Equal(t, protocol, trace.ProtocolGrpc)
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_TRACES_PROTOCOL": "http/protobuf",
		})
		defer envvars.Unset()

		env, err := config.NewFromEnv()
		require.NoError(t, err)

		var traceConfig trace.Config = config.NewTrace(env)
		protocol := traceConfig.ExporterOTLPTracesProtocol()
		assert.Equal(t, protocol, trace.ProtocolHttpProtobuf)
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_PROTOCOL":        "http/protobuf",
			"OTEL_EXPORTER_OTLP_TRACES_PROTOCOL": "grpc",
		})
		defer envvars.Unset()

		env, err := config.NewFromEnv()
		require.NoError(t, err)

		var traceConfig trace.Config = config.NewTrace(env)
		protocol := traceConfig.ExporterOTLPTracesProtocol()
		assert.Equal(t, protocol, trace.ProtocolGrpc)
	})
}

func TestTrace_ExporterOTLPTracesCompression(t *testing.T) {
	t.Run("", func(t *testing.T) {
		env, err := config.NewFromEnv()
		require.NoError(t, err)

		var traceConfig trace.Config = config.NewTrace(env)
		compression := traceConfig.ExporterOTLPTracesCompression()
		assert.Equal(t, compression, trace.CompressionGzip)
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_TRACES_COMPRESSION": "none",
		})
		defer envvars.Unset()

		env, err := config.NewFromEnv()
		require.NoError(t, err)

		var traceConfig trace.Config = config.NewTrace(env)
		compression := traceConfig.ExporterOTLPTracesCompression()
		assert.Equal(t, compression, trace.CompressionNone)
	})
}
