package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	config "github.com/yazmeyaa/virtual_ideas_book/internal/config"
	database "github.com/yazmeyaa/virtual_ideas_book/internal/database"
	logger "github.com/yazmeyaa/virtual_ideas_book/internal/logger"
	idea_repository "github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea"
	ideacomment_repository "github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea_comment"
	ideatag_repository "github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea_tag"
	ideavote_repository "github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea_vote"
	user_repository "github.com/yazmeyaa/virtual_ideas_book/internal/repository/user"
	authentication_service "github.com/yazmeyaa/virtual_ideas_book/internal/service/authentication"
	idea_service "github.com/yazmeyaa/virtual_ideas_book/internal/service/idea"
	user_service "github.com/yazmeyaa/virtual_ideas_book/internal/service/user"
	auth_controller "github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/authentication"
	idea_controller "github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/idea"
	router "github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/router"
)

func main() {
	log := logger.NewLogger()
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Error("Cannot load config")
		panic("cannot load config")
	}
	db, err := database.NewDatabaseConnection(config.DSN(), log)
	if err != nil {
		panic("cannot initialize database")
	}

	err = database.Migrate(db)
	if err != nil {
		log.Error("error while migrate database", "error", err.Error())
		return
	}

	ideaRepo := idea_repository.New(db, log)
	ideatagRepo := ideatag_repository.New(db, log)
	ideaCommentRepo := ideacomment_repository.New(db, log)
	ideaVoteRepo := ideavote_repository.New(db, log)
	userRepo := user_repository.New(db, log)
	validate := validator.New()

	ideaService := idea_service.New(
		ideaRepo,
		ideaCommentRepo,
		ideatagRepo,
		ideaVoteRepo,
		log,
	)

	authService := authentication_service.New(
		config,
	)

	userService := user_service.New(
		userRepo,
		authService,
	)

	authController := auth_controller.New(
		userService,
		authService,
		log,
		validate,
	)

	ideaController := idea_controller.New(
		ideaService,
	)

	router := router.NewRouter(
		authController,
		ideaController,
		authService,
		log,
	)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Failed to start server:", "error", err.Error())
		}
	}()

	log.Info("Server started successfull")

	<-stop
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown: %v", "error", err.Error())
	}

	log.Info("Server exited gracefully")
}
