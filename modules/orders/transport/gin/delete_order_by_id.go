package ginOrder

import (
	"first-app/common"
	"first-app/modules/orders/biz"
	"first-app/modules/orders/models"
	"first-app/modules/orders/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func DeleteOrder(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.TableCreate

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

		business := biz.NewDeleteOrderByIdBiz(store)

		if err := business.DeleteNewOrderById(c.Request.Context(),id,&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}	

		c.JSON(http.StatusOK,common.SimpleSuccessResponse(true))
	}
}