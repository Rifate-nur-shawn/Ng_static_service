package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Server ServerConfig
	DB     DBConfig
	Redis  RedisConfig
}

// ServerConfig holds the server port
type ServerConfig struct {
	GRPCPort string `mapstructure:"grpc_port"`
	HTTPPort string `mapstructure:"http_port"`
}

// DBConfig holds all database configuration
type DBConfig struct {
	URL string `mapstructure:"url"` // e.g., "postgresql://user:pass@host:port/db"
}

// RedisConfig holds the Redis connection info
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig() (*Config, error) {
	v := viper.New()

	// Set config file path
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./internal/config") // Path to config file
	v.AddConfigPath(".")                 // For running from root

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		log.Printf("Warning: could not read config file: %v. Using defaults and env vars.", err)
	}

	// Set environment variable replacer
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// Set defaults
	v.SetDefault("server.grpc_port", ":50051")
	v.SetDefault("server.http_port", ":8080")
	v.SetDefault("db.url", "postgresql://postgres:password@localhost:5432/matrimonial_db?sslmode=disable")
	v.SetDefault("redis.addr", "localhost:6379")
	v.SetDefault("redis.password", "")
	v.SetDefault("redis.db", 0)

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
