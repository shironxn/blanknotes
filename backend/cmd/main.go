package main

import (
	"github.com/shironxn/blanknotes/internal/adapter/http/handler"
	"github.com/shironxn/blanknotes/internal/adapter/http/middleware"
	"github.com/shironxn/blanknotes/internal/adapter/http/route"
	"github.com/shironxn/blanknotes/internal/adapter/repository"
	"github.com/shironxn/blanknotes/internal/config"
	"github.com/shironxn/blanknotes/internal/core/domain"
	"github.com/shironxn/blanknotes/internal/core/service"
	"github.com/shironxn/blanknotes/internal/util"

	_ "github.com/shironxn/blanknotes/docs"

	"github.com/charmbracelet/log"
)

// @title gocrud
// @version 1.0
// @description golang crud api
// @BasePath /api/v1
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := config.NewGorm(cfg).Connection()
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&domain.User{}, &domain.Note{}, &domain.RefreshToken{})

	validator, err := util.NewValidator()
	if err != nil {
		log.Fatal(err)
	}

	app := config.NewFiber()
	bcrypt := util.NewBcrypt()
	jwt := util.NewJWT(cfg)
	pagination := util.NewPagination(validator)

	userRepository := repository.NewUserRepository(db, pagination)
	userService := service.NewUserService(userRepository, bcrypt)
	userHandler := handler.NewUserHandler(userService, validator, jwt)

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository, bcrypt, jwt, cfg)
	authHandler := handler.NewAuthHandler(authService, jwt, validator, cfg)

	noteRepository := repository.NewNoteRepository(db, pagination)
	noteService := service.NewNoteService(noteRepository)
	noteHandler := handler.NewNoteHandler(noteService, validator, jwt, cfg)

	authMiddleware := middleware.NewAuthMiddleware(authService, jwt, cfg)

	initRoute := route.NewInitRoute(cfg)
	authRoute := route.NewAuthRoute(authHandler, authMiddleware)
	userRoute := route.NewUserRoute(userHandler, authMiddleware)
	noteRoute := route.NewNoteRoute(noteHandler, authMiddleware)

	initRoute.Route(app)
	authRoute.Route(app)
	userRoute.Route(app)
	noteRoute.Route(app)

	if err = app.Listen(cfg.Server.Host + ":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
