package ginMenu

import (
	"first-app/common"
	"first-app/modules/menus/biz"
	"first-app/modules/menus/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func DeleteMenu(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSQLStore(db);

		business := biz.NewDeleteMenuByIdBiz(store)

		if err := business.DeleteMenuById(c.Request.Context(),id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}	

		c.JSON(http.StatusOK,common.SimpleSuccessResponse(true))
	}
}