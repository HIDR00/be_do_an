package ginDevices

import (
	"first-app/common"
	"first-app/modules/devices/biz"
	"first-app/modules/devices/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListDevices(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Process()

		store := storage.NewSQLStore(db)

		business := biz.NewGetListDevicesBiz(store)
		result, err := business.GetNewListDevices(c.Request.Context(),&paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
