package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

type postgres_config struct{
	Database struct {
		Host						string 'json:"host"'
		Port						string  'json:"port"'
		AdminUsername						string  'json:"admin_username"'
		Adminpassword						string  'json:"admin_password"'
	}'json:database'
	
}
func main() {
	postgres_json, err := os.Open("postgres_config_file.json")
	if err != nil  {
		fmt.Println(err)
	}
	var postgres_cfg postgres_config
	:


}
