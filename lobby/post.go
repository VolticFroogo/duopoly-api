package lobby

import (
	"context"
	"fmt"
	"net/http"

	"github.com/VolticFroogo/duopoly-api/db"
	"github.com/VolticFroogo/duopoly-api/helper"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Post(c *gin.Context) {
	// Get request from JSON.
	var request struct {
		Server string `json:"server"`
	}

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("binding json: %s", err.Error()),
		})
		return
	}

	ctx := context.Background()

	lobby := Lobby{
		Server: request.Server,
	}

	// Generate ID while avoiding a conflict with two identical IDs.
	// This technically doesn't avoid all conflicts but an error would
	// require two IDs with 62^16 possible combinations generating in a ~200ms window.
	for {
		lobby.ID = helper.RandomString(IDLen)

		count, err := db.Lobbies.CountDocuments(ctx, bson.M{
			"_id": lobby.ID,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("counting lobbies with id: %s", err.Error()),
			})
			return
		}

		if count == 0 {
			break
		}
	}

	// Insert the new lobby into the DB.
	_, err = db.Lobbies.InsertOne(ctx, lobby)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("inserting lobby: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, lobby)
}
