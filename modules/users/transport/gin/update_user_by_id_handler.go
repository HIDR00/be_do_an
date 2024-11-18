package ginUser 

import (
	"first-app/common"
	"first-app/modules/users/biz"
	"first-app/modules/users/models"
	"first-app/modules/users/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func UpdateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.User

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewUpdateUserByIdBiz(store)

		if err := business.UpdateNewUserById(c.Request.Context(),id,&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}	

		c.JSON(http.StatusOK,common.SimpleSuccessResponse(true))
	}
}