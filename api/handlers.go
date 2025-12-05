package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

var DB *gorm.db

func InitDB() {
	postgres_json, err := os.Open("postgres_config.json")
	if err != nil {
		fmt.Println(err)
	}
	var postgres_cfg postgres_config

	decoder := json.NewDecoder(postgres_json)
	if err := decoder.Decode(&postgres_cfg); err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("postgres://%s:@%s:%s/%s",
		postgres_cfg.Database.Username, postgres_cfg.Database.Host, postgres_cfg.Database.Port, postgres_cfg.Database.Dbname)
	// fmt.Println(psqlInfo)
	DB, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config)
}

func CreateMetric(c *gin.Context) {
	var metric Metric
	if err := c.ShouldBindJSON(&metric); err != nil {
		ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}
	DB.Create(&metric)
	ResponseJSON(c, http.StatusCreated, "metric created succesfully", metric)

}

func getMetrics(c *gin.Context) {
	var metrics []Metric
	DB.find(&metrics)
	ResponseJSON(c, http.StatusOK, "metrics got successfully", metrics)
}
