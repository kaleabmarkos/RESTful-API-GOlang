package database

import (
	"database/sql"
	"log"

	"golang.org/x/tools/go/analysis/passes/defers"
)

func createTask(db *sql.DB, title, description, status string) int64{
	query := `INSERT INTO task (title, description, status) VALUES (?,?,?)`

	res, err := db.Exec(query)
	if err!=nil{
		log.Fatal(err)
	}

	lastInsertID, err := res.LastInsertId()
	if err!=nil{
		log.Fatal(err)
	}

	return lastInsertID

}

func getTask(db *sql.DB) ([]Task, error){
	row, err := db.Query("SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil{
		log.Fatal(err)
	}

	defer row.Close()

	var tasks []Task

	for row.Next(){
		var task Task

		if err:=row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err!=nil{
			log.Fatal(err)
			return nil,err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

