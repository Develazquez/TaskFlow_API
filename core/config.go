package core

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	JWTSecret   string
	APIPort     string
}

func LoadConfig() *Config {
	// Intentamos cargar .env, si falla puede que las variables ya estén en el entorno (ej. Docker)
	err := godotenv.Load()
	if err != nil {
		log.Println("No se encontró el archivo .env, usando variables de entorno del sistema.")
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "taskflow_db"),
		JWTSecret:  getEnv("JWT_SECRET", "defaultsecret"),
		APIPort:    getEnv("API_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
