package ginUser 

import (
	"first-app/common"
	"first-app/modules/users/biz"
	"first-app/modules/users/models"
	"first-app/modules/users/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.User

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewCreateUserBiz(store)

		if err := business.CreateNewUser(c.Request.Context(),&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}		
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}