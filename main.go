package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// user=willpelech, db=metrics, no password, localhost:5432
	dsn := "postgres://willpelech:@localhost:5432/metrics?sslmode=disable"
	fmt.Println("DSN:", dsn)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("sql.Open failed: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("db.Ping failed: %v", err)
	}

	fmt.Println("Connected successfully!")

	rows, err := db.Query("SELECT * FROM metrics")
	if err != nil {
		log.Fatalf("test query failed: %v", err)
	}
	defer rows.Close()

	fmt.Println("Test query succeeded!")
}
