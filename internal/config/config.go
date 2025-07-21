package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// DBHost       string
	// DBPort       string
	// DBUser       string
	// DBPassword   string
	// DBName       string
	// JWTSecret    string
	ServerPort   string
	ClientKey    string
	ClientSecret string
	SPSBaseURL   string
	// ProxyURL     string
	// Cors         string
	// SPSBaseURL2  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		// DBHost:       os.Getenv("DB_HOST"),
		// DBPort:       os.Getenv("DB_PORT"),
		// DBUser:       os.Getenv("DB_USER"),
		// DBPassword:   os.Getenv("DB_PASSWORD"),
		// DBName:       os.Getenv("DB_NAME"),
		// JWTSecret:    os.Getenv("JWT_SECRET"),
		ServerPort:   os.Getenv("SERVER_PORT"),
		ClientKey:    os.Getenv("CLIENT_KEY"),
		ClientSecret: os.Getenv("CLIENT_SECRET_KEY"),
		SPSBaseURL:   os.Getenv("SPS_BASE_API"),
		// ProxyURL:     os.Getenv("PROXY_URL"),
		// Cors: os.Getenv("CORS"),
		// SPSBaseURL2:  os.Getenv("SPS_BASE_API_2"),
	}
}
