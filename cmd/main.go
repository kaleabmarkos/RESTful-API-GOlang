package main

import (
	"fmt"
	"log"
	"net/http"
	"restapi-go/pkg/database"
	"restapi-go/pkg/handlers"

	"github.com/gorilla/mux"
)

func main(){
	db:= database.Connect()
	defer db.Close()

	database.CreateTable(db)

	r := mux.NewRouter()

	r.HandleFunc("/tasks", handlers.CreateTaskHandler(db)).Methods("POST")
	r.HandleFunc("/tasks", handlers.GetAllTaskHandler(db)).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.GetTaskHandler(db)).Methods("GET")
	r.HandleFunc("/tasks/{id}", handlers.UpdateTaskHandler(db)).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTaskHandler(db)).Methods("DELETE")

	taskId := database.CreateTask(db, "Test Task", "This is a test task.", "pending")
	fmt.Printf("New task created with ID: %d\n", taskId)


	tasks, _ := database.GetTask(db)
	for _,task := range tasks{
		fmt.Println("Task: ",task)
	}

	fmt.Println("Starting server at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}