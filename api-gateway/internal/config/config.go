package config

import "os"

type Config struct {
	SchoolURL  string
	StudentURL string
	AuthURL    string
	Port       string
	JWTSecret  string
}

func Load() Config {
	return Config{
		SchoolURL:  getEnv("SCHOOL_SERVICE_URL", "http://school-service:8080"),
		StudentURL: getEnv("STUDENT_SERVICE_URL", "http://student-service:8080"),
		AuthURL:    getEnv("AUTH_SERVICE_URL", "http://auth-service:8080"),
		Port:       getEnv("PORT", "8000"),
		JWTSecret:  os.Getenv("SECRET"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
