package database

import ("log"
		"database/sql"
	)

func CreateTable(db *sql.DB){
	query := `CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        description TEXT NOT NULL,
        status TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`
	
	_,err := db.Exec(query)
	if err!=nil{
		log.Fatal(err)
	}

}