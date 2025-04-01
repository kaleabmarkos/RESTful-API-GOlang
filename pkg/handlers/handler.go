package handlers

import(
	"encoding/json"
	"net/http"
	"restapi-go/pkg/database"
	"restapi-go/pkg/models"
	"database/sql"
)

func CreateTaskHandler(db *sql.DB) http.HandlerFunc{
	//handles creating a new task
}

func GetAllTaskHandler(db *sql.DB) http.HandlerFunc{
	//handles getting all task
}

func GetTaskHandler(db *sql.DB) http.HandlerFunc{
	//handles getting a task
}

func UpdateTaskHandler(db *sql.DB) http.HandlerFunc{
	//handles updating a task
}

func DeleteTaskHandler(db *sql.DB) http.HandlerFunc{
	//handles deleting a task
}

