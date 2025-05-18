package app

import (
	"context"
	"github.com/sorawaslocked/ap2final_base/pkg/logger"
	mongocfg "github.com/sorawaslocked/ap2final_base/pkg/mongo"
	grpcserver "github.com/sorawaslocked/ap2final_movie_service/internal/adapter/grpc"
	mongorepo "github.com/sorawaslocked/ap2final_movie_service/internal/adapter/mongo"
	"github.com/sorawaslocked/ap2final_movie_service/internal/config"
	"github.com/sorawaslocked/ap2final_movie_service/internal/usecase"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const serviceName = "movie service"

type App struct {
	grpcServer *grpcserver.Server
	log        *slog.Logger
}

func New(
	ctx context.Context,
	cfg *config.Config,
	log *slog.Logger,
) (*App, error) {
	const op = "App.New"

	newLog := log.With(slog.String("op", op))
	newLog.Info("starting service", slog.String("service", serviceName))

	// Mongo
	newLog.Info("connecting to mongo database", slog.String("uri", cfg.Mongo.URI))

	db, err := mongocfg.NewDB(ctx, cfg.Mongo)
	if err != nil {
		newLog.Error("error connecting to mongo database", logger.Err(err))

		return nil, err
	}

	// Repository
	movieRepo := mongorepo.NewMovie(db.Connection)

	// Use case
	movieUseCase := usecase.NewMovie(movieRepo)

	// GRPC server
	grpcServer := grpcserver.New(cfg.Server.GRPC, log, movieUseCase)

	return &App{
		grpcServer: grpcServer,
		log:        log,
	}, nil
}

func (a *App) stop() {
	a.grpcServer.Stop()
}

func (a *App) Run() {
	a.grpcServer.MustRun()

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	s := <-shutdownCh
	a.log.Info("received system shutdown signal", slog.Any("signal", s))
	a.stop()
	a.log.Info("graceful shutdown complete")
}
