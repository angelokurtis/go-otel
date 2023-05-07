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
