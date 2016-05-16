package ginyourface

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	facecontrolResponse "github.com/gobricks/facecontrol/classes/response"
)

// Facecontrol middleware checks for token is session cookie on every request
func Facecontrol() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookies := c.Request.Cookies()

		sessionCookieName := os.Getenv("FC_SESSION_COOKIE")
		facecontrolHost := os.Getenv("FC_HOST")
		loginPage := os.Getenv("FC_LOGIN_PAGE")

		if sessionCookieName == "" {
			panic("Facecontrol session cookie variable is not set")
		}

		var cookieToken string
		for _, cookie := range cookies {
			if cookie.Name == sessionCookieName {
				cookieToken = cookie.Value
				break
			}
		}

		if cookieToken == "" {
			c.Redirect(http.StatusTemporaryRedirect, loginPage)
		}

		response, err := http.Get(facecontrolHost + "/token/" + cookieToken)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			c.AbortWithStatus(http.StatusForbidden)
		}

		var responseJSON facecontrolResponse.UserResponse
		json.NewDecoder(response.Body).Decode(&responseJSON)

		c.Set("userPayload", responseJSON.User)

		c.Next()
	}
}
