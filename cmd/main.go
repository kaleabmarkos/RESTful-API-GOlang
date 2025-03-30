package main

import (
	"fmt"
	"restapi-go/pkg/database"
)

func main(){
	db:= database.Connect()
	defer db.Close()

	database.CreateTable(db)

	taskId := database.CreateTask(db, "Test Task", "This is a test task.", "pending")
	fmt.Printf("New task created with ID: %d\n", taskId)


	tasks, _ := database.GetTask(db)
	for _,task := range tasks{
		fmt.Println("Task: ",task)
	}

}