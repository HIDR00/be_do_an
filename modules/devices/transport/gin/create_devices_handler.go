package ginDevices

import (
	"first-app/common"
	"first-app/modules/devices/biz"
	"first-app/modules/devices/models"
	"first-app/modules/devices/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateDevices(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.Devices

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewCreateDevicesBiz(store)

		if err := business.CreateNewDevices(c.Request.Context(),&data,data.DeviceToken); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}		
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}