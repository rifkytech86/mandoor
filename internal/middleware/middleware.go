package middleware

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {

	protectedPaths := map[string]bool{
		"/v1/create_user": true,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := protectedPaths[r.URL.Path]; ok {
			log.Info().Msg("Middleware logic applied")
		}
		next.ServeHTTP(w, r)
	})
}
