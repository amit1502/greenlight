package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {

	cfg := config{}

	flag.IntVar(&cfg.port, "port", 8000, "Server port")
	flag.StringVar(&cfg.env, "env", "development", "Env (production | development | staging")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/healthcheck", app.healthCheckHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on port %d\n", app.config.env,
		app.config.port)

	err := srv.ListenAndServe()
	logger.Fatal(err)

}
