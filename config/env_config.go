package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	ServerAddress string
	JWTSecretKey  string
	KafkaBrokers  string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (Config, error) {
	// Load environment variables from .env file if it exists
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get values from environment variables and handle error if not found
	dbHost, err := getEnv("DB_HOST")
	if err != nil {
		return Config{}, err
	}

	dbPort, err := getEnv("DB_PORT")
	if err != nil {
		return Config{}, err
	}

	dbUser, err := getEnv("DB_USER")
	if err != nil {
		return Config{}, err
	}

	dbPassword, err := getEnv("DB_PASSWORD")
	if err != nil {
		return Config{}, err
	}

	dbName, err := getEnv("DB_NAME")
	if err != nil {
		return Config{}, err
	}

	serverAddress, err := getEnv("SERVER_ADDRESS")
	if err != nil {
		return Config{}, err
	}

	jwtSecretKey, err := getEnv("JWT_SECRET_KEY")
	if err != nil {
		return Config{}, err
	}

	kafkaBrokers, err := getEnv("KAFKA_BROKERS")
	if err != nil {
		return Config{}, err
	}

	// Return the loaded configuration
	return Config{
		DBHost:        dbHost,
		DBPort:        dbPort,
		DBUser:        dbUser,
		DBPassword:    dbPassword,
		DBName:        dbName,
		ServerAddress: serverAddress,
		JWTSecretKey:  jwtSecretKey,
		KafkaBrokers:  kafkaBrokers,
	}, nil
}

// getEnv retrieves environment variable and returns an error if it's not found
func getEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("environment variable %s not found", key)
	}
	return value, nil
}
