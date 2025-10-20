package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	database "github.com/VergilX/my-space/internal/db"
	"github.com/VergilX/my-space/internal/dblayer"
)

const version = "1.0.0"

type config struct {
	port         int
	env          string
	db           string
	traceEnabled bool
}

// dependency injection components
type application struct {
	config  config
	logger  *slog.Logger
	querier *dblayer.Queries
}

func main() {
	var cfg config

	// Adding flags
	flag.IntVar(&cfg.port, "port", 4000, "API Server Port")
	flag.StringVar(&cfg.env, "environment", "development", "Environment(development|staging|production)")
	flag.StringVar(&cfg.db, "database name", "dump.db", "Database name")
	flag.BoolVar(&cfg.traceEnabled, "trace", false, "stack trace enable for errors")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := database.New(cfg.db)
	if err != nil {
		log.Fatalf("database connection error. Error: %v", err)
		os.Exit(1)
	}
	defer db.CloseConn()

	app := application{
		config:  cfg,
		logger:  logger,
		querier: db.Queries,
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
