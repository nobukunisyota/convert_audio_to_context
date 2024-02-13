package main

import (
	"net/http"

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
	e.POST("/tasks", addTask)
	e.GET("/tasks", getTasks)
	e.PUT("/tasks", updateTask)
	e.DELETE("/tasks", deleteTask)

	return e
}
func addTask(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func getTasks(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func updateTask(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func deleteTask(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
