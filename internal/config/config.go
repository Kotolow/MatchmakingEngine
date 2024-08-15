package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port           string  `mapstructure:"SERVICE_PORT"`
	RedisAddr      string  `mapstructure:"REDIS_ADDR"`
	RedisPW        string  `mapstructure:"REDIS_PW"`
	RedisDB        int     `mapstructure:"REDIS_DB"`
	GroupSize      int     `mapstructure:"GROUP_SIZE"`
	MaxSkillDiff   float64 `mapstructure:"MAX_SKILL_DIFF"`
	MaxLatencyDiff float64 `mapstructure:"MAX_LATENCY_DIFF"`
}

var AppConfig Config

func ConfigInit() {
	var err error
	AppConfig, err = LoadConfig(".")
	if err != nil {
		fmt.Printf("error loading config: %v", err)
	}
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
