package api

import (
	"github.com/gin-gonic/gin"
	"time"
)

type JsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// so the metric class has the following pattter
type Metric struct {
	Id        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}
