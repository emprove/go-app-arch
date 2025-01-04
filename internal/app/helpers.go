package app

import (
	"fmt"
	"net/http"

	"go-app-arch/internal/infrastructure/logging"
)

func (app *Application) backgroundTask(r *http.Request, fn func() error) {
	app.WG.Add(1)

	go func() {
		defer app.WG.Done()

		defer func() {
			if err := recover(); err != nil {
				logging.LogRequestError(r, fmt.Errorf("%v", err))
			}
		}()

		if err := fn(); err != nil {
			logging.LogRequestError(r, err)
		}
	}()
}
