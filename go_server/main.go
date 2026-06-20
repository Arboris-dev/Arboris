package main

import (
	"Arboris/go_server/client/reviewer"
	"Arboris/go_server/config"
	"Arboris/go_server/webhook/api"
	"log"
	"log/slog"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func main() {

	envVar, loadErr := config.LoadEnv()

	if loadErr != nil {
		slog.Error("Unable to load the env variables", "ERROR", loadErr)
	} else {
		slog.Info("Loaded env....")
	}

	_, conn, reviewServerErr := reviewer.ConnectToPython(envVar)

	if reviewServerErr != nil {
		slog.Error("Unable to connect to python client", "ERROR", reviewServerErr)
	} else {
		slog.Info("Successfully connected to python client")
	}

	defer conn.Close()

	port := envVar.GoServer.Port
	host := envVar.GoServer.Host
	address := host + ":" + port
	log.Println("Host:Port ", address)

	client := redis.Client{}
	webhookRouter, hookRouterErr := api.NewHookRouter(*envVar, &client)

	if hookRouterErr != nil {
		slog.Error("Unable to generate webhook router", "ERROR", hookRouterErr)
	}

	server := &http.Server{
		Handler: webhookRouter,
		Addr:    address,
	}

	err := server.ListenAndServe()

	if err != nil {
		slog.Error("Server Start error", "ERROR", err)
	}

	slog.Info("Server started", "PORT", server.Addr)

}
