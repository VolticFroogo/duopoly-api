package lobby

import (
	"context"
	"fmt"
	"net/http"

	"github.com/VolticFroogo/duopoly-api/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Get(c *gin.Context) {
	// Get requested ID.
	id := c.Param("id")
	if len(id) != IDLen {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("id must have length of %d", IDLen),
		})
		return
	}

	ctx := context.Background()

	// Find lobby and return it if it exists.
	var lobby Lobby
	err := db.Lobbies.FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&lobby)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("finding lobby: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, lobby)
}
