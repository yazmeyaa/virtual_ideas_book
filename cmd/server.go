package main

import (
	common "github.com/yazmeyaa/virtual_ideas_book/internal/common"
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

	ideaService := idea_service.New(
		ideaRepo,
		ideaCommentRepo,
		ideatagRepo,
		ideaVoteRepo,
		log,
	)

	user, err := userRepo.GetByUsername("zmeya")

	if err != nil {
		log.Warn("Error while create user", "error", err.Error())
		return
	}

	log.Info("Created user info",
		"username", user.Username,
	)

	ideaService.CreateIdea(idea_repository.CreateIdeaDTO{
		Title:       "Test",
		Description: "Test Description",
		Content:     "Hello world",
		UserID:      user.ID,
	})

	log.Info("Successfull started server")

	list, _ := ideaService.GetIdeasList(idea_repository.GetIdeaListOptions{
		ListOptions: common.ListOptions{
			Limit:  10,
			Offset: 0,
		},
	})

	authService := authentication_service.New(config)
	token, err := authService.CreateToken(user)

	if err != nil {
		log.Error("ERROR WHILE GET TOKEN", "error", err.Error())
		return
	}
	log.Info("Token successfull generated", "token", token)

	isTokenValid, err := authService.ValidateToken(token)
	if err != nil {
		log.Error("ERROR WHILE VALIDATE TOKEN", "error", err.Error())
		return
	}

	log.Info("Is token valid", "valid", isTokenValid)

	for _, idea := range list {
		log.Info("Idea info",
			"id", idea.ID,
			// "title", idea.Title,
			// "content", idea.Content,
			// "tags", idea.Tags,
		)
	}
}
