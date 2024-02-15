package controller

import (
	"net/http"
	"sample-api/model"
	"strconv"

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
	model.DB.First(&task, id)
	return c.JSON(http.StatusOK, task)
}
func UpdateTask(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}
	task.ID = id
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
	return c.JSON(http.StatusOK, task)
}
