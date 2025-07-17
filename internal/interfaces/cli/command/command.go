package command

import (
	"context"
	"go-app-arch/internal/app"
)

type Command interface {
	Run(ctx context.Context, args []string) *error
}

type SitemapGen struct {
	app *app.Application
}

func NewSitemapGenCmd(app *app.Application) *SitemapGen {
	return &SitemapGen{app: app}
}
