package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyas-muhammad/bookstore_oauth_api/src/domain/access_token"
	"github.com/ilyas-muhammad/bookstore_oauth_api/src/http"
	"github.com/ilyas-muhammad/bookstore_oauth_api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	if err := router.Run(":8080"); err != nil {
		panic(err.Error())
	}
}
