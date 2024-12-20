package main

import (
	"github.com/aventhis/go-url-shortener/internal/config"
	"github.com/aventhis/go-url-shortener/internal/lib/logger/sl"
	"github.com/aventhis/go-url-shortener/internal/storage/sqlite"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// init config
	cfg := config.MustLoad()

	// init logger
	log := setupLogger(cfg.Env)

	log.Info("starting url shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	// init storage: sqlite
	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to initialize storage", sl.Err(err))
		os.Exit(1)
	}
	defer storage.Close()

	//_, err = storage.SaveURL("https://www.youtube.com", "youtube")
	//if err != nil {
	//	log.Error("failed to save url", sl.Err(err))
	//}
	//
	//url, err := storage.GetURL("youtube")
	//if err != nil {
	//	log.Error("failed to get url", sl.Err(err))
	//}
	//log.Info("url for alias youtube", slog.String("url", url))

	//err = storage.DeleteURL("youtube")
	//if err != nil {
	//	log.Error("failed to delete url", sl.Err(err))
	//}
	//
	//log.Info("Delete url", slog.String("url", url))

	// TODO: init router: chi, "chi render"

	// TODO: run server:

}

func setupLogger(env string) *slog.Logger {
	var handler slog.Handler
	switch env {
	case envLocal:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case envDev:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case envProd:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}

	return slog.New(handler)
}
