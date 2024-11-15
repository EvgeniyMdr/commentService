package config

import (
	"os"
)

var mainConfig Config

func init() {
	mainConfig = Config{
		dbSettings: DbConfig{
			User:     getEnv("POSTGRES_USER", "CPostgres"),
			Password: getEnv("POSTGRES_PASSWORD", "CPostQwerty"),
			Name:     getEnv("POSTGRES_DB", "commentdb"),
			Host:     getEnv("POSTGRES_HOST", "postgresdb"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
		},
		httpSettings: HttpSettings{
			Url:  getEnv("HTTP_URL", "0.0.0.0"),
			Port: getEnv("HTTP_PORT", "8081"),
		},
		grpcSettings: GrpcSettings{
			Host:    getEnv("GRPC_HOST", "0.0.0.0"),
			Port:    getEnv("GRPC_PORT", "44030"),
			TimeOut: getEnv("GRPC_TIMEOUT", "5s"),
		},
	}
}

type ServiceConfig struct {
	cfg *Config
}

func (sc *ServiceConfig) GetDbSettings() DbConfig {
	return sc.cfg.dbSettings
}

func (sc *ServiceConfig) GetHttpSettings() HttpSettings {
	return sc.cfg.httpSettings
}

func (sc *ServiceConfig) GetGRPCSettings() GrpcSettings {
	return sc.cfg.grpcSettings
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{
		cfg: &mainConfig,
	}
}

type Config struct {
	dbSettings   DbConfig
	httpSettings HttpSettings
	grpcSettings GrpcSettings
}

// TODO: Узнать нужны ли теги без библиотеки для считывания env переменных по тегам?
type DbConfig struct {
	User     string `env:"POSTGRES_USER" env-default:"CPostgres"`
	Password string `env:"POSTGRES_PASSWORD" env-default:"CPostQwerty"`
	Name     string `env:"POSTGRES_DB" env-default:"commentdb"`
	Host     string `env:"POSTGRES_HOST" env-default:"postgresdb"`
	Port     string `env:"POSTGRES_PORT" env-default:"5432"`
}

type HttpSettings struct {
	Url  string `env:"HTTP_URL" env-default:"0.0.0.0"`
	Port string `env:"HTTP_PORT" env-default:"8081"`
}

type GrpcSettings struct {
	Host    string `env:"GRPC_HOST" env-default:"0.0.0.0"`
	Port    string `env:"GRPC_POST" env-default:"44030"`
	TimeOut string `env:"GRPC_TIMEOUT" env-default:"5s"`
}

func (hs *HttpSettings) GetAddress() string {
	return hs.Url + ":" + hs.Port
}
