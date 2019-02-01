package custom_middleware

import "github.com/gin-gonic/gin"

func ServerHeader(c *gin.Context) {
	// shortcut for c.Writer.Header().Set("Server", "Some-Play-Server")
	c.Header("Content-Encoding", "gzip")
	c.Header("Content-Type", "application/json; charset=utf-8")
}
