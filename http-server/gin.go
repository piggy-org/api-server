package http_server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GinServer gin http handler
type GinServer struct {
	engine *gin.Engine
}

// NewGinServer create a gin http http-server
func NewGinServer() Server {
	engine := gin.New()
	server := &GinServer{
		engine: engine,
	}
	engine.GET(healthPath, server.health)
	return server
}

// Run a gin http-server
func (h *GinServer) Run() error {
	err := h.engine.Run(defaultPort)
	return err
}

// Stop a gin http-server
func (h *GinServer) Stop() error { return nil }

// health is a handler for gin http-server health check
func (h *GinServer) health(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Response{
		Code: ok,
		Data: `{"msg": "ok"}`,
	})
}
