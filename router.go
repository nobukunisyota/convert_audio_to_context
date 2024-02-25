package main

import (
	"sample-api/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
	// Create a new instance of Echo
	e := echo.New()

	// Setting up middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	// Define the route
	api := e.Group("/api/v1")
	api.POST("/tasks", controller.AddTask)
	api.GET("/tasks/:id", controller.GetTaskByID)
	api.GET("/tasks", controller.GetTasks)
	api.PUT("/tasks/:id", controller.UpdateTask)
	api.DELETE("/tasks/:id", controller.DeleteTask)
	// api.POST("/users", controller.Login)

	return e
}
