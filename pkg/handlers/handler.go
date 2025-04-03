package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"restapi-go/pkg/database"
	"restapi-go/pkg/models"
	"strconv"
	"strings"
	"github.com/gorilla/mux"
)

func CreateTaskHandler(db *sql.DB) http.HandlerFunc{
	//handles creating a new task and storing it the database
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(task.Title)=="" || strings.TrimSpace(task.Description)==""{
			http.Error(w, "Task and Description are required", http.StatusBadRequest)
		}
		validStatus := map[string]bool{"pending":true, "completed":true}
		if !validStatus[task.Status]{
			http.Error(w, "Invalid Status", http.StatusBadRequest)
		}

		taskId := database.CreateTask(db, task.Title, task.Description, task.Status)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{"id":taskId})
	}

}

func GetAllTaskHandler(db *sql.DB) http.HandlerFunc{
	//handles getting all task
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := database.GetTask(db)
		if err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)

	}
}


func GetTaskHandler(db *sql.DB) http.HandlerFunc{
	//handles getting a task
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err:=strconv.Atoi(idStr)
		if err!=nil{
			http.Error(w, "Invalid Id", http.StatusBadRequest)
			return
		}

		tasks, err:= database.GetTask(db)
		if err!=nil{
			http.Error(w,"Task not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks[id])
	}
}

func UpdateTaskHandler(db *sql.DB) http.HandlerFunc{
	//handles updating a task
	return func(w http.ResponseWriter, r *http.Request) {
		idStr :=mux.Vars(r)["id"]
		idint, err := strconv.Atoi(idStr)
		if err != nil{
			http.Error(w,"Invalid Id", http.StatusBadRequest)
		}

		var task models.Task
		
		err = json.NewDecoder(r.Body).Decode(&task)
		if err!=nil{
			http.Error(w,err.Error(), http.StatusBadRequest)
			return
		}

		err = database.UpdateTask(db, idint, task.Title, task.Description, task.Status)
		if err !=nil{
			http.Error(w, "Error updating the task", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status":"updated"})

	}
}

func DeleteTaskHandler(db *sql.DB) http.HandlerFunc{
	//handles deleting a task
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		idint, err := strconv.Atoi(idStr)
		if err!=nil{
			http.Error(w,"Invalid ID", http.StatusBadRequest)
			return
		}

		err = database.DeleteTask(db, idint)
		if err!=nil{
			http.Error(w,err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status":"deleted"})

	}
}

