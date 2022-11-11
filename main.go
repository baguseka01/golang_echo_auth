package main

import (
	"github.com/baguseka01/golang_echo_mongodb/configs"
	"github.com/baguseka01/golang_echo_mongodb/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Initialize database
	configs.ConnectDB()

	// Initialize routes
	routes.UserRouter(e)

	// Port
	e.Logger.Fatal(e.Start(":6000"))
}
