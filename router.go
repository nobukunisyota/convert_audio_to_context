package main

import (
	"fmt"
	"net/http"
	"sample-api/model"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
	// Create a new instance of Echo
	e := echo.New()

	// Setting up middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	fmt.Printf("e: %T", e)
	// Define the route
	e.File("/", "index.html")
	e.POST("/tasks", addTask)
	e.GET("/tasks/:id", getTaskByID)
	e.GET("/tasks", getTasks)
	e.PUT("/tasks/:id", updateTask)
	e.DELETE("/tasks/:id", deleteTask)

	return e
}
func addTask(c echo.Context) error {
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	model.DB.Create(&task)
	return c.JSON(http.StatusCreated, task)
}
func getTasks(c echo.Context) error {
	tasks := []model.Task{}
	model.DB.Find(&tasks)
	return c.JSON(http.StatusOK, tasks)
}
func getTaskByID(c echo.Context) error {
	id := c.Param("id")
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	model.DB.First(&task, id)
	return c.JSON(http.StatusOK, task)
}
func updateTask(c echo.Context) error {
	// id := c.Param("id")
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	model.DB.Save(&task)
	return c.JSON(http.StatusOK, task)
}
func deleteTask(c echo.Context) error {
	id := c.Param("id")
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	delete(task, id)
	return c.JSON(http.StatusOK, task)
}
