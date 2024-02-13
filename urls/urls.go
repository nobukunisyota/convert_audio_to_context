package urls

import "github.com/labstack/echo"

func AddTask(c echo.Context) {
	println("addTask")
}
func GetTasks(c echo.Context) {
	println("getTasks")
}
func UpdateTask(c echo.Context) {
	println("updateTask")
}
func DeleteTask(c echo.Context) {
	println("deleteTask")
}
