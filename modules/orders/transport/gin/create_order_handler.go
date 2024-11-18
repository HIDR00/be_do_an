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

func CreateOrder(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var data models.TableCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewCreateOrderBiz(store)

		if err := business.CreateNewOrder(c.Request.Context(),&data,id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}		
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}