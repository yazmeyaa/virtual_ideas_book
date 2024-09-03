package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	authentication_service "github.com/yazmeyaa/virtual_ideas_book/internal/service/authentication"
	authentication_handler "github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/authentication"
)

type UserIdKeyType string

const (
	USER_ID_KEY UserIdKeyType = "UserID"
)

func AuthMiddleware(next http.Handler, authService authentication_service.AuthenticationService, log *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeaders := r.Header.Get(authentication_handler.AUTH_TOKEN_COOKIE_NAME)

		if authHeaders == "" {
			log.Warn("Unauthorized access attempt",
				"reason", "missing Authorization header",
				"path", r.URL.Path,
				"ip", r.RemoteAddr,
			)
			http.Error(w, "Unauthorized - missed auth token", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(authHeaders, " ")
		if len(headerParts) != 2 {
			log.Info("Malformed Authorization header",
				"header", authHeaders,
				"path", r.URL.Path,
				"ip", r.RemoteAddr,
			)
			http.Error(w, "Wrong authentication headers signature", http.StatusUnauthorized)
			return
		}

		tokenString := headerParts[1]
		isTokenValid, err := authService.ValidateToken(tokenString)
		if !isTokenValid || err != nil {
			log.Info("Invalid token",
				"token", tokenString,
				"path", r.URL.Path,
				"ip", r.RemoteAddr,
			)
			http.Error(w, "Unauthorized - invalid or missing token.", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), USER_ID_KEY, "")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
