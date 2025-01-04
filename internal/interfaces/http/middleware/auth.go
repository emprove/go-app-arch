package middleware

import (
	"net/http"

	"go-app-arch/internal/domain/service"
)

func Authenticate(sUser service.UserServiceInterface) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			token := req.Header.Get("Authorization")
			if token == "" {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			_, err := sUser.FindOneByToken(token)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, req)
		})
	}
}
