package ginOrder

import (
	"first-app/common"
	"first-app/modules/orders/biz"
	"first-app/modules/orders/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetOrderById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewGetOrderByIdBiz(store)

		data,err := business.GetNewOrderById(c.Request.Context(),id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		
		c.JSON(http.StatusOK, common.BasicSuccessResponse(data))
	}
}