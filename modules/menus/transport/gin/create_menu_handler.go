package ginMenu

import (
	"first-app/common"
	"first-app/modules/menus/biz"
	"first-app/modules/menus/models"
	"first-app/modules/menus/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateMenu(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.Menu

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewCreateMenuBiz(store)

		if err := business.CreateNewMenu(c.Request.Context(),&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}		
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}