package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kaniek99/graph_miner/app"
)

const port = 8080

func main() {
	srv := app.NewServer(port)

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("Failed to start server: %v\n", err)
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	sig := <-shutdown
	log.Printf("Received signal: %s\n", sig)
	if err := srv.Stop(); err != nil {
		log.Fatalf("Failed to shutdown server: %v\n", err)
	}
}
