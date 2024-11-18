package ginTable

import (
	"first-app/common"
	"first-app/modules/tables/biz"
	"first-app/modules/tables/models"
	"first-app/modules/tables/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTable(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.Table

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewCreateTableBiz(store)

		if err := business.CreateNewTable(c.Request.Context(),&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}		
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}