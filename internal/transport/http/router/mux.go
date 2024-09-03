package router

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yazmeyaa/virtual_ideas_book/internal/service/authentication"
	authController "github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/authentication"
	ideaController "github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/idea"
	"github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/middleware"
)

func NewRouter(
	authController authController.AuthenticationHandler,
	ideaController ideaController.IdeaHandler,
	authService authentication.AuthenticationService,
	log *slog.Logger,
) *mux.Router {
	router := mux.NewRouter()

	v1 := router.PathPrefix("/v1").Subrouter()

	authRouter := v1.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authController.Login).Methods("POST")
	authRouter.HandleFunc("/register", authController.Register).Methods("POST")

	ideaRouter := v1.PathPrefix("/idea").Subrouter()
	ideaRouter.Use(func(h http.Handler) http.Handler {
		return middleware.AuthMiddleware(h, authService, log)
	})
	ideaRouter.HandleFunc("/{id}", ideaController.GetById).Methods("GET")
	ideaRouter.HandleFunc("/:id", ideaController.Update).Methods("PATCH")
	ideaRouter.HandleFunc("/:id", ideaController.Delete).Methods("DELETE")
	ideaRouter.HandleFunc("/", ideaController.Get).Methods("GET")

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("404 not found.",
			"path", r.URL.Path,
			"ip", r.RemoteAddr,
		)
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	return router
}
