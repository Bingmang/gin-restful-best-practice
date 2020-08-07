package routes

import (
	"gin-restful-best-practice/controllers"
	"gin-restful-best-practice/middlewares"
)

func init() {

	// unauthorized
	unauthorizedAPI := engine.Group("api")
	{
		unauthorizedAPI.POST("/users/login", controllers.Login)
		unauthorizedAPI.POST("/users", controllers.CreateUser)
	}

	// authorized with jwt
	authorizedAPI := engine.Group("api")
	authorizedAPI.Use(middlewares.AuthenticateJWT())
	{
		authorizedAPI.GET("/models", controllers.FetchModelList)
		authorizedAPI.POST("/models", controllers.CreateModel)

		authorizedAPI.GET("/users", controllers.FetchUserList)
		authorizedAPI.GET("/users/:id", controllers.FetchUserByID)
	}
}
