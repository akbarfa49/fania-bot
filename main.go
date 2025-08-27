package main

import (
	"context"
	"fania-bot/core"
	"fania-bot/transport/http"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	ctx := context.Background()
	core := core.New()
	go core.RunStreamNotifier()

	httpServer := &http.Server{HostPort: []string{":8080"}, CoreService: core}
	httpServer.ServeNBHTTP()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	httpServer.Shutdown(ctx)
}
