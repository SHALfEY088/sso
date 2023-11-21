// `cmd/sso/main.go`
package main

import (
	"github.com/SHALfEY088/sso/internal/config"
	"github.com/SHALfEY088/sso/internal/lib/logger/handlers/slogpretty"
	"github.com/SHALfEY088/sso/internal/lib/logger/sl"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	//application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	//
	//application.GRPCSrv.MustRun()

	log.Info("starting application", slog.Any("config", cfg))

	var err error

	if err != nil {
		log.Error("error message", sl.Err(err))
	}

	// TODO: инициализировать логгер

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
