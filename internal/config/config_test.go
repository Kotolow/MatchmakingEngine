package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	t.Run("Valid config", func(t *testing.T) {
		config, err := LoadConfig("../..")
		assert.NoError(t, err)

		assert.Equal(t, ":8080", config.Port)
		assert.Equal(t, "172.20.0.2:6379", config.RedisAddr)
		assert.Equal(t, "", config.RedisPW)
		assert.Equal(t, 0, config.RedisDB)
		assert.Equal(t, 5, config.GroupSize)
		assert.Equal(t, float64(500), config.MaxSkillDiff)
		assert.Equal(t, float64(90), config.MaxLatencyDiff)
	})
}
