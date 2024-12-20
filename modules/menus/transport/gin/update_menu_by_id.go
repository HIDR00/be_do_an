package ginMenu

import (
	"first-app/common"
	"first-app/modules/menus/biz"
	"first-app/modules/menus/models"
	"first-app/modules/menus/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func UpdateMenu(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.Menu

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

		business := biz.NewUpdateMenuByIdBiz(store)

		if err := business.UpdateNewMenuById(c.Request.Context(),id,&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}	

		c.JSON(http.StatusOK,common.SimpleSuccessResponse(true))
	}
}