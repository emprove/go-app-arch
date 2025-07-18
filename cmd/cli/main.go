package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"runtime/debug"
	"time"

	"go-app-arch/internal/app"
	"go-app-arch/internal/interfaces/cli/command"
)

var CmdMap = map[string]func(app *app.Application) command.Command{
	"sitemap-gen": func(app *app.Application) command.Command {
		return command.NewSitemapGenCmd(app)
	},
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	app, err := app.NewApp()
	if err != nil {
		slog.Error(err.Error(), "trace", string(debug.Stack()))
		os.Exit(1)
	}

	defer app.DB.Close()

	cmdName := flag.String("c", "", "command name")
	flag.Parse()

	if cmdConstructor, exists := CmdMap[*cmdName]; exists {
		cmd := cmdConstructor(app)
		if err := cmd.Run(ctx, flag.Args()); err != nil {
			err := *err
			slog.Error(err.Error(), "trace", string(debug.Stack()))
			os.Exit(1)
		}
	} else {
		slog.Error("Unknown command", "command", *cmdName)
		os.Exit(1)
	}

	os.Exit(0)
}
