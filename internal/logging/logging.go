package logging

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

func LogRequestError(r *http.Request, err error) {
	var (
		message = err.Error()
		method  = r.Method
		url     = r.URL.String()
		trace   = string(debug.Stack())
	)

	requestAttrs := slog.Group("request", "method", method, "url", url)
	slog.Error(message, requestAttrs, "trace", trace)
}
