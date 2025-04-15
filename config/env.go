package config

import (
	"github.com/joho/godotenv"
	"github.com/caarlos0/env/v11"
)

type Env struct{

	SERVER_PORT  string `env:"SERVER_PORT"`
	DB_HOST     string `env:"DB_HOST"`
	DB_NAME     string `env:"DB_NAME"`
	DB_USER     string `env:"DB_USER"`
	DB_PASSWORD string `env:"DB_PASSWORD"`
	DB_SSLMODE  string `env:"DB_SSLMODE"`

}

func GetEnv() (*Env, error){
	
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	config := &Env{}

	err =  env.Parse(config)

	if err != nil {
		return nil, err
	}

	return config, nil

}
