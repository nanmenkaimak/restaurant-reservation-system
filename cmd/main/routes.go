package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nanmenkaimak/restaurant-reservation-system/internal/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func routes() {
	router := gin.Default()

	router.NoRoute(func(ctx *gin.Context) { // check for 404
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Page not found",
		})
	})

	router.GET("/", handlers.Repo.Home)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user := router.Group("/user")
	{
		user.POST("/signup", handlers.Repo.SignUp)
		user.POST("/login", handlers.Repo.Login)
	}

	restaurant := router.Group("/restaurant", Auth())
	{
		restaurant.GET("", handlers.Repo.Rests)
		restaurant.GET("/:rest_id", handlers.Repo.SingleRest)
		restaurant.POST("/new", handlers.Repo.AddRest)
		restaurant.POST("/:rest_id/table", handlers.Repo.AddTable)
		restaurant.GET("/:rest_id/table", handlers.Repo.ShowAllTableOfRest)
		restaurant.GET("/owner/:owner_id", handlers.Repo.AllOwnerRests)
		restaurant.POST("/reservations/:seat_id", handlers.Repo.ReserveTable)
		restaurant.GET("/owner/reservations", handlers.Repo.ShowAllReservationsOfRest)
		restaurant.POST("/:rest_id/menu", handlers.Repo.AddFoodToMenu)
		restaurant.GET("/:rest_id/menu", handlers.Repo.ShowMenuOfRest)
	}

	_ = router.Run(portNumber)
}
