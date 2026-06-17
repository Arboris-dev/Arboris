package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoServer struct {
		host string
		port string
	}
	PyServer struct {
		host string
		port string
	}
}

func GetConfig() Config {
	err := godotenv.Load()

	if err != nil {
		slog.Info("Couldn't initialize godotenv. Skipping the loading.....")
	}

	var config Config

	config.GoServer.host = os.Getenv("GO_SERVER_HOST")
	config.GoServer.port = os.Getenv("GO_SERVER_PORT")

	config.PyServer.host = os.Getenv("PYTHON_SERVER_HOST")
	config.PyServer.port = os.Getenv("PYTHON_SERVER_PORT")

	return config
}
