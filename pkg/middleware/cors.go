package middleware

import (
	"net/http"
	"regexp"
)

type CorsConfig struct {
	AllowedOrigin string
}

func allowedOrigin(cfg *CorsConfig, origin string) bool {
	if cfg.AllowedOrigin == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(cfg.AllowedOrigin, origin); matched {
		return true
	}
	return false
}

func Cors(cfg *CorsConfig) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if allowedOrigin(cfg, r.Header.Get("Origin")) {
				w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
			}
			if r.Method == "OPTIONS" {
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
