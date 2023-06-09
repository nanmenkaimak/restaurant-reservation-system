package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/handlers"
	"net/http"
)

func routes() http.Handler {
	router := gin.Default()

	router.NoRoute(func(ctx *gin.Context) { // check for 404
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	})

	router.GET("/", handlers.Repo.Home)

	user := router.Group("/user")
	{
		user.POST("/signup", handlers.Repo.SignUp)
		user.POST("/login", handlers.Repo.Login)
	}

	restaurant := router.Group("/restaurant", Auth())
	{
		restaurant.GET("/", handlers.Repo.AllRests)
		restaurant.POST("/owner", handlers.Repo.AddRest)
	}

	return router
}
