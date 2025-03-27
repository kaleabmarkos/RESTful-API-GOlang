package models

import (
	"fmt"
	"net/http"
	"time"
)

type Task struct{
	ID				int			`json:"id"`
	Title			string 		`json:"title"`
	Description		string		`json:"description"`
	Status			string		`json:"status"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
}

type User struct{
	ID				int			`json:"id"`
	UserName		string 		`json:"username"`
	Email			string		`json:"email"`
	PassWord		string		`json:"password"`
}

