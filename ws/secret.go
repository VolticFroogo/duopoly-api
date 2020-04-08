package ws

import (
	"net/http"
	"time"

	"github.com/VolticFroogo/duopoly-api/helper"
	"github.com/gin-gonic/gin"
)

const (
	secretLen      = 32
	secretName     = "secret"
	secretDuration = time.Hour * 24 * 365
)

func getSecret(c *gin.Context, h *http.Header) (secret string, err error) {
	// Get the secret from the secret cookie.
	secret, err = c.Cookie(secretName)
	if err != nil {
		// If there was an error which wasn't a no cookie error, it's probably bad; return.
		if err != http.ErrNoCookie {
			return
		}

		// If we reach here, there was just no cookie found; let's create one.
		// This is expected behaviour and not an error; nullify the error.
		err = nil

		// Generate the new secret as a random string.
		secret = helper.RandomString(secretLen)

		// Create a cookie with all of the relevant parameters.
		cookie := http.Cookie{
			Name:     secretName,
			Value:    secret,
			Path:     "/",
			Expires:  time.Now().Add(secretDuration),
			Secure:   false, // TODO: change to secure in production.
			HttpOnly: true,
		}

		// Add the cookie to the HTTP header as Set-Cookie.
		h.Add("Set-Cookie", cookie.String())
	}

	return
}
