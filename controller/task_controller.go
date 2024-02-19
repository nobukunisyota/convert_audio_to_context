package controller

import (
	"net/http"
	"sample-api/model"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func AddTask(c echo.Context) error {
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	model.DB.Create(&task)
	return c.JSON(http.StatusCreated, task)
}
func GetTasks(c echo.Context) error {
	tasks := []model.Task{}
	model.DB.Find(&tasks)
	return c.JSON(http.StatusOK, tasks)
}
func GetTaskByID(c echo.Context) error {
	id := c.Param("id")
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	if err := model.DB.First(&task, id); err.Error != nil {
		return c.JSON(http.StatusNotFound, "Task not found")
	}
	return c.JSON(http.StatusOK, task)
}
func UpdateTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	updateAt := time.Now()
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	task.ID = id
	task.UpdatedAt = updateAt
	model.DB.Save(&task)
	return c.JSON(http.StatusOK, task)
}
func DeleteTask(c echo.Context) error {
	id := c.Param("id")
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	model.DB.Delete(&task, id)
	return c.JSON(http.StatusOK, "Task deleted")
}
