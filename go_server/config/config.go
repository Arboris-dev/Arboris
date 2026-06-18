package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoServer struct {
		Host string
		Port string
	}
	PyServer struct {
		Host string
		Port string
	}
	WebHook struct {
		PayloadMaxSize int
		Burst          int
	}
}

func GetConfig() Config {
	err := godotenv.Load()

	if err != nil {
		slog.Info("Couldn't initialize godotenv. Skipping the loading.....")
	}

	var config Config

	config.GoServer.Host = os.Getenv("GO_SERVER_HOST")
	config.GoServer.Port = os.Getenv("GO_SERVER_PORT")

	config.PyServer.Host = os.Getenv("PYTHON_SERVER_HOST")
	config.PyServer.Port = os.Getenv("PYTHON_SERVER_PORT")

	config.WebHook.Burst = os.Getenv("")
	return config
}
