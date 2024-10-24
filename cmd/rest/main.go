package main

import (
	"log/slog"
	"os"
	"runtime/debug"

	"go-app-arch/internal/app"
	"go-app-arch/internal/rest"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		slog.Error(err.Error(), "trace", string(debug.Stack()))
		os.Exit(1)
	}

	defer app.DB.Close()

	router := rest.NewRouter(app.Cfg, app.DS, app.DB)
	rest.ServeHTTP(app.Cfg, router)

	app.WG.Wait()
}
