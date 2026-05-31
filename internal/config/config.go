package config

import "github.com/spf13/viper"

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	JWTSecret  string
}

func LoadConfig() (*Config, error){
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		AppPort:    viper.GetString("APP_PORT"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBSSLMode:  viper.GetString("DB_SSLMODE"),
		JWTSecret:  viper.GetString("JWT_SECRET"),
	},nil
}