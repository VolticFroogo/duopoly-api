package lobby

import "github.com/gin-gonic/gin"

const (
	IDLen = 16
)

type Lobby struct {
	ID     string `bson:"_id"    json:"id"`
	Server string `bson:"server" json:"server"`
}

func Init(r *gin.Engine) {
	// Register endpoints with Gin.
	r.GET("/lobby/:id", Get)
	r.POST("/lobby", Post)
}
