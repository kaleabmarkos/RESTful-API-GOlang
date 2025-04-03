package database

import (
	"database/sql"
	"log"
	"restapi-go/pkg/models"
)

func CreateTask(db *sql.DB, title, description, status string) int64{
	query := `INSERT INTO tasks (title, description, status) VALUES (?,?,?)`

	res, err := db.Exec(query, title, description, status)
	if err!=nil{
		log.Fatal(err)
	}

	lastInsertID, err := res.LastInsertId()
	if err!=nil{
		log.Fatal(err)
	}

	return lastInsertID

}

func GetTask(db *sql.DB) ([]models.Task, error){
	row, err := db.Query("SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil{
		log.Fatal(err)
	}

	defer row.Close()

	var tasks []models.Task

	for row.Next(){
		var task models.Task

		if err:=row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err!=nil{
			log.Fatal(err)
			return nil,err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func UpdateTask(db *sql.DB, id int, title, description, status string) error {
	query := `UPDATE tasks SET title = ?, description = ?, status = ?, update_at = CURENT_TIMESTAMP WHERE ID = ?`
	_, err := db.Exec(query, title, description, status, id)
	if err!=nil{
		log.Fatal(err)
		return err	
	}

	return nil

}

func DeleteTask(db *sql.DB, id int) error {
	query :=`DELETE FROM tasks WHERE id = ?`
	_,err := db.Exec(query, id)
	if err != nil{
		log.Fatal(err)
		return err
	}

	return nil
}


