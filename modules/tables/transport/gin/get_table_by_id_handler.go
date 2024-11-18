package ginTable

import (
	"first-app/common"
	"first-app/modules/tables/biz"
	"first-app/modules/tables/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTableById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewGetTableByIdBiz(store)

		data,err := business.GetNewTableById(c.Request.Context(),id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}