package server

import "github.com/gin-gonic/gin"

// router path
const (
	healthPath = "/health"
)

type HealthHandler interface {
	health(c *gin.Context)
}

// retCode service return code
type retCode int

const (
	ok retCode = iota
)

// Response common response struct
type Response struct {
	Code retCode     `json:"code"`
	Data interface{} `json:"data"`
}

const (
	defaultPort = ":8080"
)
