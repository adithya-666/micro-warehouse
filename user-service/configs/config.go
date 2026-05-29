package configs

import (
	"github.com/spf13/viper"
)

type App struct {
	AppPort string `json:"app_port"`
	AppEnv  string `json:"app_env"`
}

type Database struct {
	Host             string `json:"_host"`
	Port             string `json:"_port"`
	User             string `json:"_user"`
	Password         string `json:"_password"`
	DBName           string `json:"db_name"`
	DBMaxOpenConns   int    `json:"db_max_open_conns"`
	DBMaxIdleConns   int    `json:"db_max_idle_conns"`
	DBMaxConnIdleSec int    `json:"db_max_conn_idle_sec"`
}

type Redis struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type RabbitMQ struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Supabase struct {
	URL    string `json:"url"`
	APIKey string `json:"api_key"`
	Bucket string `json:"bucket"`
}

type Config struct {
	App      App      `json:"app"`
	Database Database `json:"database"`
	Redis    Redis    `json:"redis"`
	RabbitMQ RabbitMQ `json:"rabbitmq"`
	Supabase Supabase `json:"supabase"`
}

func NewConfig() *Config {
	return &Config{
		App: App{
			AppPort: viper.GetString("APP_PORT"),
			AppEnv:  viper.GetString("APP_ENV"),
		},
		Database: Database{
			Host:             viper.GetString("DATABASE_HOST"),
			Port:             viper.GetString("DATABASE_PORT"),
			User:             viper.GetString("DATABASE_USER"),
			Password:         viper.GetString("DATABASE_PASSWORD"),
			DBName:           viper.GetString("DATABASE_NAME"),
			DBMaxOpenConns:   viper.GetInt("DATABASE_MAX_OPEN_CONNECTION"),
			DBMaxIdleConns:   viper.GetInt("DATABASE_MAX_IDLE_CONNECTION"),
			DBMaxConnIdleSec: viper.GetInt("DATABASE_MAX_CONN_IDLE_SEC"),
		},
		Redis: Redis{
			Host: viper.GetString("REDIS_HOST"),
			Port: viper.GetString("REDIS_PORT"),
		},
		RabbitMQ: RabbitMQ{
			Host:     viper.GetString("RABBITMQ_HOST"),
			Port:     viper.GetString("RABBITMQ_PORT"),
			User:     viper.GetString("RABBITMQ_USER"),
			Password: viper.GetString("RABBITMQ_PASSWORD"),
		},
		Supabase: Supabase{
			URL:    viper.GetString("SUPABASE_URL"),
			APIKey: viper.GetString("SUPABASE_KEY"),
			Bucket: viper.GetString("SUPABASE_BUCKET"),
		},
	}

}