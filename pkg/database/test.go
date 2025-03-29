package database

import ("database/sql"
		"log"
	)

func createTask(db *sql.DB, title, description, status string) int64{
	query := `INSERT INTO task (title, description, status) VALUES (?,?,?)`

	res, err := db.Exec(query)
	if err!=nil{
		log.Fatal(err)
	}

	lastInsertID, err := res.LastInsertId()
	if err!=nil{
		log.Fatal(err)
	}

	return lastInsertID

}