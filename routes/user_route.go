package routes

import (
	"github.com/baguseka01/golang_echo_mongodb/controllers"
	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Echo) {
	e.POST("/user", controllers.CreateUser)
	e.GET("/user/:userId", controllers.DetailUser)
	e.PUT("/user/:userId", controllers.UpdateUser)
	e.DELETE("/user/:userId", controllers.DeleteUser)
	e.GET("/users", controllers.GetAllUser)
}