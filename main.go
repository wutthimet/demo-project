package main

import (
	"demo-project/controller"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	config "github.com/spf13/viper"
	"github.com/tylerb/graceful"
	"os"
	"time"
)

func init() {
	env := os.Args[1]

	// API Start
	fmt.Println("Server start running environment configuration : ", env)

	config.SetConfigFile("configs/" + env + ".yml")
	if err := config.ReadInConfig(); err != nil {
		fmt.Println("Fatal error env config (file): ", err)
	}
}

func main() {

	// Initial ECHO Framework
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	// Routes
	r := e.Group(config.GetString("service.endpoint"))
	r.GET("/monitoring", new(controller.Monitoring).Monitoring)

	// Start Server, Graceful Shutdown with in 5sec.
	e.Server.Addr = ":" + config.GetString("service.port")
	graceful.ListenAndServe(e.Server, 5*time.Second)
}
