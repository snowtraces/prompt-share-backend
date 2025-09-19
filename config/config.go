package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Addr string `mapstructure:"addr"`
}

type JWTConfig struct {
	Secret      string `mapstructure:"secret"`
	ExpireHours int    `mapstructure:"expire_hours"`
}

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	Path   string `mapstructure:"path"`
}

type LocalStorageConfig struct {
	BasePath string `mapstructure:"base_path"`
}

type StorageConfig struct {
	Local LocalStorageConfig `mapstructure:"local"`
}

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Database DatabaseConfig `mapstructure:"database"`
	Storage  StorageConfig  `mapstructure:"storage"`
}

var Cfg *Config

func Load() {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("read config error:", err)
	}
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatal("unmarshal config error:", err)
	}
	Cfg = &cfg
}
