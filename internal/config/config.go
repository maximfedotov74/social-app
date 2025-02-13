package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Database struct {
	Port         int
	Host         string
	User         string
	Password     string
	DatabaseName string
}

type Minio struct {
	Port     int
	Host     string
	User     string
	Password string
}

type Jwt struct {
	RefreshTokenExpires string
	AccessTokenExpires  string
	RefreshTokenSecret  string
	AccessTokenSecret   string
}

type HTTP struct {
	Port int
}

type SMTP struct {
	Port int
	Host string
	Key  string
	Mail string
}

type Redis struct {
	Port     int
	Host     string
	Password string
}

type Config struct {
	Database *Database
	Minio    *Minio
	Jwt      *Jwt
	HTTP     *HTTP
	SMTP     *SMTP
	Redis    *Redis
}

var once sync.Once

var config *Config

func MustLoadConfig() *Config {
	once.Do(func() {
		app_env, ok := os.LookupEnv("APP_ENV")
		if !ok {
			log.Fatal("App environment is not exists! Please pass the application variable before launching.")
		}
		err := godotenv.Load("." + app_env + ".env")

		if err != nil {
			log.Fatalf("Error when load env configuration! Cause: %s", err.Error())
		}

		config = &Config{
			Database: &Database{
				Port:         getEnvAsInt("DATABASE_PORT"),
				Host:         getEnv("DATABASE_HOST"),
				User:         getEnv("DATABASE_USER"),
				Password:     getEnv("DATABASE_PASSWORD"),
				DatabaseName: getEnv("DATABASE_NAME"),
			},
			Minio: &Minio{
				Port:     getEnvAsInt("MINIO_PORT"),
				Host:     getEnv("MINIO_HOST"),
				User:     getEnv("MINIO_USER"),
				Password: getEnv("MINIO_PASSWORD"),
			},
			Jwt: &Jwt{
				RefreshTokenExpires: getEnv("REFRESH_TOKEN_EXP"),
				AccessTokenExpires:  getEnv("ACCESS_TOKEN_EXP"),
				RefreshTokenSecret:  getEnv("REFRESH_TOKEN_SECRET"),
				AccessTokenSecret:   getEnv("ACCESS_TOKEN_SECRET"),
			},
			HTTP: &HTTP{
				Port: getEnvAsInt("HTTP_PORT"),
			},
			SMTP: &SMTP{
				Port: getEnvAsInt("SMTP_PORT"),
				Host: getEnv("SMTP_HOST"),
				Key:  getEnv("SMTP_KEY"),
				Mail: getEnv("SMTP_MAIl"),
			},
			Redis: &Redis{
				Port:     getEnvAsInt("REDIS_PORT"),
				Host:     getEnv("REDIS_HOST"),
				Password: getEnv("REDIS_PASSWORD"),
			},
		}
	})
	return config
}
