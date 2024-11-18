package ginTable

import (
	"first-app/common"
	"first-app/modules/tables/biz"
	"first-app/modules/tables/models"
	"first-app/modules/tables/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func UpdateTable(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.Table

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

		business := biz.NewUpdateTableByIdBiz(store)

		if err := business.UpdateNewTableById(c.Request.Context(),id,&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}	

		c.JSON(http.StatusOK,common.SimpleSuccessResponse(true))
	}
}