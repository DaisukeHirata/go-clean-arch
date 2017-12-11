package main

import (
	httpDeliver "github.com/bxcodec/go-clean-arch/article/delivery/http"
	articleRepo "github.com/bxcodec/go-clean-arch/article/repository/mysql"
	articleUcase "github.com/bxcodec/go-clean-arch/article/usecase"
	"github.com/bxcodec/go-clean-arch/config/middleware"
	"github.com/labstack/echo"
)

func (a *App) initializeRouter() {
	a.Router = echo.New()
	middL := middleware.InitMiddleware()
	a.Router.Use(middL.CORS)
}

func (a *App) setupHandler() {
	// article
	ar := articleRepo.NewMysqlArticleRepository(a.DB)
	au := articleUcase.NewArticleUsecase(ar)
	httpDeliver.NewArticleHttpHandler(a.Router, au)

	// job

	// product
}
