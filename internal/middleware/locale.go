package middleware

import (
	"net/http"
	"slices"

	"go-app-arch/internal/config"
)

func Locale(availableLocalesIso []string, ds *config.DynamicState) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			lang := req.Header.Get("Accept-Language")
			if lang != "" && slices.Contains(availableLocalesIso, lang) {
				ds.Locale = lang
			}

			next.ServeHTTP(w, req)
		})
	}
}
