package router

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yazmeyaa/virtual_ideas_book/internal/service/authentication"
	authController "github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/authentication"
	ideaController "github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/idea"
	ideaTagController "github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/idea_tag"
	"github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/middleware"
)

func NewRouter(
	authController authController.AuthenticationHandler,
	ideaController ideaController.IdeaHandler,
	ideaTagController ideaTagController.IdeaTagHTTPHanlder,
	authService authentication.AuthenticationService,
	log *slog.Logger,
) *mux.Router {
	router := mux.NewRouter()

	v1 := router.PathPrefix("/v1").Subrouter()

	// Auth routes
	authRouter := v1.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", authController.Login).Methods("POST")
	authRouter.HandleFunc("/register", authController.Register).Methods("POST")

	// Public idea routes
	publicIdeaRouter := v1.PathPrefix("/idea").Subrouter()
	publicIdeaRouter.StrictSlash(true)
	publicIdeaRouter.HandleFunc("/", ideaController.Get).Methods("GET")
	publicIdeaRouter.HandleFunc("/{id}", ideaController.GetById).Methods("GET")

	// Protected idea routes (with AuthMiddleware)
	protectedIdeaRouter := v1.PathPrefix("/idea").Subrouter()
	protectedIdeaRouter.Use(func(h http.Handler) http.Handler {
		return middleware.AuthMiddleware(h, authService, log)
	})
	protectedIdeaRouter.HandleFunc("/", ideaController.Create).Methods("POST")
	protectedIdeaRouter.HandleFunc("/{id}", ideaController.Update).Methods("PATCH")
	protectedIdeaRouter.HandleFunc("/{id}", ideaController.Delete).Methods("DELETE")

	ideaTagRouter := v1.PathPrefix("/tag").Subrouter()
	publicIdeaRouter.StrictSlash(true)
	ideaTagRouter.HandleFunc("{id}", ideaTagController.GetTagById).Methods("GET")
	ideaTagRouter.HandleFunc("", ideaTagController.GetTags).Methods("GET")
	ideaTagRouter.HandleFunc("", ideaTagController.Create).Methods("POST")
	// ideaTagRouter.HandleFunc("/{id}", ideaTagController.Delete).Methods("DELETE")
	ideaTagRouter.HandleFunc("{id}", ideaTagController.Update).Methods("PATCH")

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("404 not found.",
			"path", r.URL.Path,
			"method", r.Method,
			"ip", r.RemoteAddr,
		)
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	return router
}
