package model

import (
	"github.com/capelaum/go-task-api/views"
)

//ReadAll - list all registers from table tasks
func ReadAll() ([]views.PostRequest, error) {

	rows, err := con.Query("SELECT * FROM tasks")
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	// create array of tasks
	tasks := []views.PostRequest{}

	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.ID, &data.Name, &data.Task)
		tasks = append(tasks, data)
	}

	return tasks, nil
}

//ReadByName Search for name in the table tasks
func ReadByName(name string) ([]views.PostRequest, error) {

	rows, err := con.Query("SELECT * FROM tasks WHERE name=?", name)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	// create array of tasks
	tasks := []views.PostRequest{}

	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.ID, &data.Name, &data.Task)
		tasks = append(tasks, data)
	}

	return tasks, nil
}
