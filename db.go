package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"os"
)

type postgres_config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Username string `json:"admin_username"`
		Password string `json:"admin_password"`
		Dbname   string `json:"dbname"`
	} `json:"database"`
}

func main() {
	postgres_json, err := os.Open("postgres_config.json")
	if err != nil {
		fmt.Println(err)
	}
	var postgres_cfg postgres_config

	decoder := json.NewDecoder(postgres_json)
	if err := decoder.Decode(&postgres_cfg); err != nil {
		panic(err)
	}

	// fmt.Println("successfully read JSON")
	// fmt.Println("Host", postgres_cfg.Database.Host)
	//  "postgres://%s"//
	//  urlExample := "postgres://username:password@localhost:5432/database_name"
	psqlInfo := fmt.Sprintf("postgres://%s:@%s:%s/%s",
		postgres_cfg.Database.Username, postgres_cfg.Database.Host, postgres_cfg.Database.Port, postgres_cfg.Database.Dbname)

	fmt.Println(psqlInfo)
	db, err := sql.Open("pgx", psqlInfo)
	db.Ping()
	if err != nil {
		panic(err)
	}

	query := "SELECT * FROM metrics"

	test, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	defer test.Close()
}
