package main

import (
	"fmt"
	"restapi-go/pkg/database"
)

func main(){
	db:= database.Connect()
	defer db.Close()

	taskId := database.createTask(db, "Test Task", "This is a test task.", "pending")
	fmt.Printf("New task created with ID: %d\n", taskId)


	tasks, _ := database.getTask(db)
	for _,task := range tasks{
		fmt.Println("Task: ",task)
	}

}