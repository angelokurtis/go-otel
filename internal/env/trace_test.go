package env_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/angelokurtis/go-starter-otel/_test"
	"github.com/angelokurtis/go-starter-otel/internal/env"
	"github.com/angelokurtis/go-starter-otel/internal/trace"
)

func TestTrace_TracesExporter(t *testing.T) {
	t.Run("", func(t *testing.T) {
		otel, err := env.LookupOTel()
		require.NoError(t, err)

		tr := env.ToTrace(otel)
		assert.Equal(t, []trace.Exporter{trace.ExporterOtlp}, tr.Exporters)
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_TRACES_EXPORTER": "logging,zipkin",
		})
		defer envvars.Unset()

		otel, err := env.LookupOTel()
		require.NoError(t, err)

		tr := env.ToTrace(otel)
		assert.Equal(t, []trace.Exporter{trace.ExporterLogging, trace.ExporterZipkin}, tr.Exporters)
	})
}

func TestTrace_ExporterOTLPTracesEndpoint(t *testing.T) {
	t.Run("", func(t *testing.T) {
		otel, err := env.LookupOTel()
		require.NoError(t, err)

		tr := env.ToTrace(otel)
		assert.Equal(t, "http://localhost:4317", tr.Endpoint.String())
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_TRACES_ENDPOINT": "https://otel.com:4317",
		})
		defer envvars.Unset()

		otel, err := env.LookupOTel()
		require.NoError(t, err)

		tr := env.ToTrace(otel)
		assert.Equal(t, "https://otel.com:4317", tr.Endpoint.String())
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_ENDPOINT":        "https://otel.com.br:4317",
			"OTEL_EXPORTER_OTLP_TRACES_ENDPOINT": "https://otel.com:4317",
		})
		defer envvars.Unset()

		otel, err := env.LookupOTel()
		require.NoError(t, err)

		tr := env.ToTrace(otel)
		assert.Equal(t, "https://otel.com:4317", tr.Endpoint.String())
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_ENDPOINT": "https://otel.com.br:4317",
		})
		defer envvars.Unset()

		otel, err := env.LookupOTel()
		require.NoError(t, err)

		tr := env.ToTrace(otel)
		assert.Equal(t, "https://otel.com.br:4317", tr.Endpoint.String())
	})
}

func TestTrace_ExporterOTLPTracesTimeout(t *testing.T) {
	t.Run("", func(t *testing.T) {
		otel, err := env.LookupOTel()
		require.NoError(t, err)

		tr := env.ToTrace(otel)

		var oneSecond int64 = 1000

		assert.Equal(t, oneSecond*10, tr.Timeout.Milliseconds())
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_TIMEOUT": "1m",
		})
		defer envvars.Unset()

		otel, err := env.LookupOTel()
		require.NoError(t, err)

		tr := env.ToTrace(otel)

		var oneMinute int64 = 60000

		assert.Equal(t, oneMinute*1, tr.Timeout.Milliseconds())
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_TIMEOUT":        "10m",
			"OTEL_EXPORTER_OTLP_TRACES_TIMEOUT": "1h",
		})
		defer envvars.Unset()

		otel, err := env.LookupOTel()
		require.NoError(t, err)

		tr := env.ToTrace(otel)

		var oneMinute int64 = 60000

		assert.Equal(t, oneMinute*60, tr.Timeout.Milliseconds())
	})
}

func TestTrace_ExporterOTLPTracesProtocol(t *testing.T) {
	t.Run("", func(t *testing.T) {
		otel, err := env.LookupOTel()
		require.NoError(t, err)

		tr := env.ToTrace(otel)
		assert.Equal(t, trace.ProtocolGrpc, tr.Protocol)
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_TRACES_PROTOCOL": "http/protobuf",
		})
		defer envvars.Unset()

		otel, err := env.LookupOTel()
		require.NoError(t, err)

		tr := env.ToTrace(otel)
		assert.Equal(t, trace.ProtocolHttpProtobuf, tr.Protocol)
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_PROTOCOL":        "http/protobuf",
			"OTEL_EXPORTER_OTLP_TRACES_PROTOCOL": "grpc",
		})
		defer envvars.Unset()

		otel, err := env.LookupOTel()
		require.NoError(t, err)

		tr := env.ToTrace(otel)
		assert.Equal(t, trace.ProtocolGrpc, tr.Protocol)
	})
}

func TestTrace_ExporterOTLPTracesCompression(t *testing.T) {
	t.Run("", func(t *testing.T) {
		otel, err := env.LookupOTel()
		require.NoError(t, err)

		tr := env.ToTrace(otel)
		assert.Equal(t, trace.CompressionGzip, tr.Compression)
	})
	t.Run("", func(t *testing.T) {
		// set the environment variables and ensure that the environment variable is cleaned up after the test
		envvars := _test.SetEnvironmentVariables(map[string]string{
			"OTEL_EXPORTER_OTLP_TRACES_COMPRESSION": "none",
		})
		defer envvars.Unset()

		otel, err := env.LookupOTel()
		require.NoError(t, err)

		tr := env.ToTrace(otel)
		assert.Equal(t, trace.CompressionNone, tr.Compression)
	})
}
