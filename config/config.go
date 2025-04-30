package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sorawaslocked/ap2final_movie_service/pkg/mongo"
	"time"
)

type (
	Config struct {
		MongoConfig mongo.Config `yaml:"mongo"`
		Server      Server       `yaml:"server"`

		Version string `yaml:"version" env:"VERSION"`
	}

	Server struct {
		GRPCServer GRPCServer `yaml:"gRPCServer"`
	}

	GRPCServer struct {
		Port                  int16         `yaml:"port" env:"GRPC_SERVER_PORT"`
		MaxReceiveMsgSize     int           `yaml:"maxReceiveMsgSize" env:"GRPC_MAX_RECEIVE_MSG_SIZE" env-default:"12"`
		MaxConnectionAge      time.Duration `yaml:"maxConnectionAge" env:"GRPC_MAX_CONNECTION_AGE" env-default:"30s"`
		MaxConnectionAgeGrace time.Duration `yaml:"maxConnectionAgeGrace" env:"GRPC_MAX_CONNECTION_AGE_GRACE" env-default:"10s"`
	}
)

func New() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error creating config: %w", err)
	}

	return &cfg, nil
}
