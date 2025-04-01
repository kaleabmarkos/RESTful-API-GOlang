package main

import (
	"fmt"
	"restapi-go/pkg/database"
	"restapi-go/pkg/handlers"
	"github.com/gorilla/mux"
	"log"
)

func main(){
	db:= database.Connect()
	defer db.Close()

	database.CreateTable(db)

	r := mux.NewRouter()

	r.HandleFunc("/tasks", handlers.CreateTaskHandler(db)).Methods("POST")

	taskId := database.CreateTask(db, "Test Task", "This is a test task.", "pending")
	fmt.Printf("New task created with ID: %d\n", taskId)


	tasks, _ := database.GetTask(db)
	for _,task := range tasks{
		fmt.Println("Task: ",task)
	}

}