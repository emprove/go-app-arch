package rest

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-app-arch/internal/config"
)

const (
	defaultIdleTimeout    = time.Minute
	defaultReadTimeout    = 5 * time.Second
	defaultWriteTimeout   = 10 * time.Second
	defaultShutdownPeriod = 30 * time.Second
)

func ServeHTTP(cfg *config.Cfg, router http.Handler) error {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.GetHttpPort()),
		Handler:      router,
		ErrorLog:     slog.NewLogLogger(slog.Default().Handler(), slog.LevelWarn),
		IdleTimeout:  defaultIdleTimeout,
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	shutdownErrorChan := make(chan error)

	go func() {
		quitChan := make(chan os.Signal, 1)
		signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
		<-quitChan

		ctx, cancel := context.WithTimeout(context.Background(), defaultShutdownPeriod)
		defer cancel()

		shutdownErrorChan <- server.Shutdown(ctx)
	}()

	slog.Info("starting server", slog.Group("server", "addr", server.Addr))

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownErrorChan
	if err != nil {
		return err
	}

	slog.Info("stopped server", slog.Group("server", "addr", server.Addr))

	return nil
}
