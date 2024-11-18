package ginUser

import (
	"context"
	"first-app/common"
	"first-app/modules/users/biz"
	"first-app/modules/users/models"
	"first-app/modules/users/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

const SCOPES = "https://www.googleapis.com/auth/firebase.messaging"

func AuthenUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.User

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewAuthenUserBiz(store)

		result, err := business.AuthenNewUser(c.Request.Context(), &data)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		//test firebase accesstoken
		token, err := getAccessToken()
		if err != nil {
			log.Fatalf("Failed to get access token: %v", err)
		}

		c.JSON(http.StatusOK, common.AuthenSuccessResponse(result, token))
	}
}

func getAccessToken() (string, error) {
	ctx := context.Background()
	// Get credentials
	creds, err := google.FindDefaultCredentials(ctx, "https://www.googleapis.com/auth/firebase.messaging")
	if err != nil {
		fmt.Print("failed to create credentials:", err)
	}

	// Generate a valid access token
	tokenSource := creds.TokenSource
	token, err := tokenSource.Token()
	if err != nil {
		fmt.Print("failed to retrieve access token:", err)
	}
	return token.AccessToken, nil
}
