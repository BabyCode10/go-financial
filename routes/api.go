package routes

import (
	"fmt"
	"go-financial/app/config"
	"go-financial/app/controllers"
	"go-financial/app/middlewares"
	"go-financial/app/services"

	"github.com/gin-gonic/gin"
)

func Api() {
	env, _ := config.Load()
	conn, err := config.Database()

	if err != nil {
		panic("Failed to load database.")
	}

	r := gin.Default()
	{
		route := r.Group("api/v1/auth")
		{
			userService := services.User(conn)
			authController := controllers.Auth(userService)

			route.POST("/register", authController.Register)
			route.POST("/login", authController.Login)
			route.POST("/me", middlewares.Auth(), authController.Me)
			route.POST("/logout", middlewares.Auth(), authController.Logout)
		}

		route = r.Group("api/v1/category")
		{
			categoryService := services.Category(conn)
			categoryController := controllers.Category(categoryService)

			route.GET("", middlewares.Auth(), categoryController.Index)
			route.POST("", middlewares.Auth(), categoryController.Create)
			route.GET(":category_id", middlewares.Auth(), categoryController.Show)
			route.PUT(":category_id", middlewares.Auth(), categoryController.Update)
			route.DELETE(":category_id", middlewares.Auth(), categoryController.Delete)
		}

		route = r.Group("api/v1/transaction")
		{
			transactionService := services.Transaction(conn)
			transactionController := controllers.Transaction(transactionService)

			route.GET("", middlewares.Auth(), transactionController.Index)
			route.POST("", middlewares.Auth(), transactionController.Create)
			route.GET(":transaction_id", middlewares.Auth(), transactionController.Show)
			route.PUT(":transaction_id", middlewares.Auth(), transactionController.Update)
			route.DELETE(":transaction_id", middlewares.Auth(), transactionController.Delete)
		}
	}

	err = r.Run(fmt.Sprintf(":%s", env.AppPort))

	if err != nil {
		panic("Failed to start server.")
	}
}
