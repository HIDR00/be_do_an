package ginOrder

import (
	"first-app/common"
	"first-app/modules/orders/biz"
	"first-app/modules/orders/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTotalOrder(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {


		store := storage.NewSQLStore(db)

		business := biz.NewGetTotalOrderByIdBiz(store)

		data,err := business.GetNewTotalOrder(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}
		
		c.JSON(http.StatusOK, common.BasicSuccessResponse(data))
	}
}