package api

import "github.com/gin-gonic/gin"

type JsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Metric struct {
}
