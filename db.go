package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// import (
//
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/jackc/pgx/v5"
//	"log"
//
// )
type postgres_config struct {
	Database struct {
		Host          string `json:"host"`
		Port          string `json:"port"`
		AdminUsername string `json:"admin_username"`
		Adminpassword string `json:"admin_password"`
		Dbname        string `json:"dbname"`
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
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		postgres_cfg.Database.host, postgres_cfg.Database.port, postgres_cfg.Database.user, postgres_cfg.Database.password, postgres_cfg.Database.Dbname)
	sql.Open("postgres", psqlInfo)

}
