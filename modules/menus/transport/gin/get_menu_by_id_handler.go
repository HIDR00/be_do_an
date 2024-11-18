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

func GetMenuById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewGetMenuByIdBiz(store)

		data,err := business.GetNewMenuById(c.Request.Context(),id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}