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
		restaurant.GET("", handlers.Repo.AllRests)
		//restaurant.GET("/:rest_id:table_id", handlers.Repo.ShowAllTableOfRest)
		restaurant.POST("/new", handlers.Repo.AddRest)
		restaurant.POST("/:rest_id/table", handlers.Repo.AddTable)
		//restaurant.GET("/", handlers.Repo.AllOwnerRest)
	}

	return router
}
