package ws

import (
	"fmt"
	"net/http"

	"github.com/VolticFroogo/duopoly-api/message"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO: actually check origin in production.
		return true
	},
	EnableCompression: true,
}

type requestGreeting struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func Handle(c *gin.Context) {
	// Declare headers to send with the WS upgrade.
	h := http.Header{}

	// Get the secret, or create one and set it in the headers.
	secret, err := getSecret(c, &h)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("getting secret: %s", err.Error()),
		})
		return
	}

	// Upgrade the HTTP connection to a WebSocket.
	conn, err := upgrader.Upgrade(c.Writer, c.Request, h)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("upgrading connection: %s", err.Error()),
		})
		return
	}

	// Free connection resources on function return.
	defer conn.Close()

	// Handle the initial greeting to find / create game and get player ID.
	g, playerID, err := greeting(conn, secret)
	if err != nil {
		// If we get an error from this function, it will be handled besides closing the conn.
		return
	}

	// Call player left once this thread dies (which should happen on a disconnection).
	defer g.PlayerLeft(playerID)

	for {
		// Declare a message to read into.
		msg := message.Message{}

		// Read the message.
		err = conn.ReadJSON(&msg)
		if err != nil {
			// Exit if there was a fatal connection error.
			if websocket.IsCloseError(err) || websocket.IsUnexpectedCloseError(err) {
				return
			}

			// Write the error to the client.
			_ = conn.WriteJSON(message.Message{
				Type: message.ResponseError,
				Data: fmt.Sprintf("reading json request: %s", err.Error()),
			})
			continue
		}

		// Pass the message on for the game to process.
		g.Message(msg, playerID)
	}
}
