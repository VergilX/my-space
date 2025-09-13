package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/VergilX/my-space/internal/db"
	database "github.com/VergilX/my-space/internal/db"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

// create application struct
type application struct {
	config config
	logger *slog.Logger
	db     *db.DB
}

func main() {
	var cfg config

	// Adding flags
	flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
	flag.StringVar(&cfg.env, "environment", "development", "Environment(development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// init db
	dsn := "file:locked.sqlite?cache=shared"
	db, err := database.New(dsn)
	if err != nil {
		log.Fatal("database connection error")
		os.Exit(1)
	}

	app := application{
		config: cfg,
		logger: logger,
		db:     db,
	}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
