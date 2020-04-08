package ws

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	// Register the endpoint with Gin.
	r.GET("/ws", Handle)
}
