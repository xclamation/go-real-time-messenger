package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// AppConfig stores application-level configurations
type Config struct {
	appEnv     string
	dbHost     string
	dbPort     int
	dbUser     string
	dbPassword string
	dbName     string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file in development
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Parse environment variables
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}
	return &Config{
		appEnv:     getEnv("APP_ENV", "development"),
		dbHost:     getEnv("DB_HOST", "localhost"),
		dbPort:     dbPort,
		dbUser:     getEnv("DB_USER", ""),
		dbPassword: getEnv("DB_PASSWORD", ""),
		dbName:     getEnv("DB_NAME", "messenger_db"),
	}
}

// Helper function to read an environment variable or return a default value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

// Getters for non-sensitive fields
func (c *Config) AppEnv() string {
	return c.appEnv
}

func (c *Config) DBHost() string {
	return c.dbHost
}

func (c *Config) DBPort() int {
	return c.dbPort
}

func (c *Config) DBName() string {
	return c.dbName
}

// Method to retrieve sensitive data (if absolutely needed) should be carefully managed
func (c *Config) DBUser() string {
	return c.dbUser
}

func (c *Config) DBPassword() string {
	return c.dbPassword
}
