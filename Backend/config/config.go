package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMODE bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env variables: ", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service Name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http Port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRECT_KEY")
	if jwtSecretKey == "" {
		fmt.Println("Jwt secret key is required")
		os.Exit(1)
	}

	db_host := os.Getenv("DB_HOST")
	if db_host == "" {
		fmt.Println("DB Host is required")
		os.Exit(1)
	}

	db_port := os.Getenv("DB_PORT")
	if db_port == "" {
		fmt.Println("DB Port is required")
		os.Exit(1)
	}

	db_prt, err := strconv.ParseInt(db_port, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	db_name := os.Getenv("DB_NAME")
	if db_name == "" {
		fmt.Println("DB Name is required")
		os.Exit(1)
	}

	db_user := os.Getenv("DB_USER")
	if db_user == "" {
		fmt.Println("DB Name is required")
		os.Exit(1)
	}

	db_pass := os.Getenv("DB_PASSWORD")
	if db_pass == "" {
		fmt.Println("DB Password is required")
		os.Exit(1)
	}

	db_enableSslMode := os.Getenv("DB_ENABLE_SSL_MODE")
	db_enableSSlMode, err := strconv.ParseBool(db_enableSslMode)
	if err != nil {
		fmt.Println("DB enable ssl mode is required")
		os.Exit(1)
	}

	dbConfig := &DBConfig{
		Host:          db_host,
		Port:          int(db_prt),
		Name:          db_name,
		User:          db_user,
		Password:      db_pass,
		EnableSSLMODE: db_enableSSlMode,
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}
}

func GetConfig() *Config { // singleton design pattern
	if configurations == nil {
		loadConfig()
	}

	return configurations
}
