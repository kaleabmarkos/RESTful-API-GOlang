package handlers

import(
	"encoding/json"
	"net/http"
	"restapi-go/pkg/database"
	"restapi-go/pkg/models"
)

func CreateTaskHandler(db *sql.DB) http.HandlerFunc{
	//handles creating a new task
}