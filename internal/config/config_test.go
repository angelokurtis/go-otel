package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/angelokurtis/go-starter-otel/internal/config"
)

func TestNewFromEnv(t *testing.T) {
	cfg, err := config.NewFromEnv()
	require.NoError(t, err)

	assert.NotEmpty(t, cfg)
}
