package model

import "fmt"

type Task struct {
	UID     int    `json:"uid"`
	ID      int    `json:"id" gorm:"praimaly_key"`
	Name    string `json:"name"`
	Context string `json:"completed"`
}
type Tasks []Task

func createTask(todo *Task) {
	db.Create(todo)
}
func getTasks(t *Task) Tasks {
	var tasks Tasks
	db.Where(t).Find(&tasks)
	return tasks
}
func UpdateTodo(t *Task) error {
	// rows := db.Model(t).Update(map[string]interface{}{
	// 	"name":      t.Name,
	// 	"completed": t.Context,
	// }).RowsAffected
	// if rows == 0 {
	// 	return fmt.Errorf("Could not find Todo (%v) to update", t)
	// }
	return nil
}
func deleteTask(t *Task) error {
	if rows := db.Where(t).Delete(&Task{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to delete", t)
	}
	return nil
}
