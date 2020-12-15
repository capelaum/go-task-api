package model

import (
	"database/sql"
	"fmt"
)

//CreateDatabase creates task_database database and table tasks
func CreateDatabase(db *sql.DB) {
	exec(db, "CREATE DATABASE IF NOT EXISTS task_database")
	exec(db, "USE task_database")
	exec(db, "DROP TABLE IF EXISTS tasks")
	exec(db, `CREATE TABLE tasks (
							id integer auto_increment,
							name varchar(100),
							task varchar(100),
							PRIMARY KEY (id)
						)`)
	fmt.Println("Created Database task_database and table tarefas reseted!")
}

//InsertTask inserts a task into tasks table
func InsertTask(id int, name, task string) error {
	insertQuery, err := con.Query("INSERT INTO tasks(id, name, task) VALUES(?, ?, ?)", id, name, task)
	defer insertQuery.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//DeleteTask deletes a task from tasks table
func DeleteTask(name string) error {
	deleteQuery, err := con.Query("DELETE FROM tasks WHERE name=?", name)
	defer deleteQuery.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
