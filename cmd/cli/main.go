package main

import (
	"flag"
	"log/slog"
	"os"
	"runtime/debug"

	"go-app-arch/internal/app"
	"go-app-arch/internal/command"
)

var CmdMap = map[string]func(app *app.Application) command.Command{
	"sitemap-gen": func(app *app.Application) command.Command {
		return command.NewSitemapGenCmd(app)
	},
}

func main() {
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
		if err := cmd.Run(flag.Args()); err != nil {
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
