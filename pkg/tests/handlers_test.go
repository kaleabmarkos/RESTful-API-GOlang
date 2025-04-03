package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restapi-go/pkg/database"
	"restapi-go/pkg/handlers"
	"restapi-go/pkg/models"
	"testing"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func testCreateTask(t *testing.T){
	db := database.Connect()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/tasks", handlers.CreateTaskHandler(db)).Methods("POST")

	newTask := models.Task{
		Title: "NEW_TITLE",
		Description: "TESTING THE CREATETASK HANDLER",
		Status: "pending",
	}

	data, err := json.Marshal(newTask)
	if err!=nil{
		t.Fatalf("ERROR marshaling the task ",err)
	}

	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(data))
	if err!=nil{
		t.Fatalf("ERROR sending request ",err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err!=nil{
		t.Fatalf("ERROR decoding response ",err)
	}

	assert.NotNil(t, response["id"])
}