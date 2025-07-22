package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Conectar() {
	connStr := "user=postgres password=1234 dbname=students_db sslmode=disable port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Conectado ao banco com sucesso!")
	DB = db
}
