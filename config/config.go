package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
	"user-service/src/common/redis"
	"user-service/src/models"
)

const (
	ApplicationPort        = "port"
	PostGreSQLHOST         = "postgresql_host"
	PostGreSQLPort         = "postgresql_port"
	PostGreSQLUsername     = "postgresql_username"
	PostGreSQLPassword     = "postgresql_password"
	PostGreSQLDatabase     = "postgresql_database"
	PostGreSQLSSLMode      = "postgresql_sslmode"
	PostGreSQLMaxIdleConns = "postgresql_max_idle_conns"
	PostGreSQLMaxOpenConns = "postgresql_max_open_conns"
	RedisUrl               = "redis_url"
	RedisPort              = "redis_port"
	RedisPassword          = "redis_password"
)

// Init configurations
func Init() {
	viper.AutomaticEnv()
	viper.AddConfigPath("./config/app/")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error read in config .env")
	}
}

// Database returns database settings
func DatabaseConfig() *models.Database {
	return &models.Database{
		Host:         viper.GetString(PostGreSQLHOST),
		Port:         viper.GetInt(PostGreSQLPort),
		Username:     viper.GetString(PostGreSQLUsername),
		Password:     viper.GetString(PostGreSQLPassword),
		Database:     viper.GetString(PostGreSQLDatabase),
		SSLMode:      viper.GetString(PostGreSQLSSLMode),
		MaxIdleConns: viper.GetInt(PostGreSQLMaxIdleConns),
		MaxOpenConns: viper.GetInt(PostGreSQLMaxOpenConns),
	}
}

func AppPort() string {
	return ":" + viper.GetString(ApplicationPort)
}

func RedisConfig() *redis.Config {
	return &redis.Config{
		Host:     viper.GetString(RedisUrl),
		Port:     viper.GetInt64(RedisPort),
		Password: viper.GetString(RedisPassword),
	}
}

func GetUserServicePort() int {
	port := viper.GetString(ApplicationPort)
	p, _ := strconv.Atoi(port[1:len(port)])
	return p
}
