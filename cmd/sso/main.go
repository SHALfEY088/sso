// `cmd/sso/main.go`
package main

import (
	"github.com/SHALfEY088/sso/internal/config"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	// ...

	// TODO: инициализировать логгер

	// TODO: инициализировать приложение (app)

	// TODO: запустить gRPC-сервер приложения
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
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
