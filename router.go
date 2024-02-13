package main

import (
	"sample-api/urls"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
	// Create a new instance of Echo
	e := echo.New()

	// Setting up middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define the route
	e.File("/", "index.html")
	e.POST("/tasks", urls.AddTask)
	e.GET("/tasks", urls.GetTasks)
	e.PUT("/tasks", urls.UpdateTask)
	e.DELETE("/tasks", urls.DeleteTask)

	return e
}
