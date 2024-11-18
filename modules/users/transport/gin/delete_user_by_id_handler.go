package ginUser 

import (
	"first-app/common"
	"first-app/modules/users/biz"
	"first-app/modules/users/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func DeleteUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		store := storage.NewSQLStore(db);

		business := biz.NewDeleteUserByIdBiz(store)

		if err := business.DeleteNewUserById(c.Request.Context(),id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}	

		c.JSON(http.StatusOK,common.SimpleSuccessResponse(true))
	}
}