package ginEmployeeHandleOrder

import (
	"first-app/common"
	"first-app/modules/employee_handle_order/biz"
	"first-app/modules/employee_handle_order/models"
	"first-app/modules/employee_handle_order/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateEmployeeHandleOrder(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data models.EmployeeHandleOrder

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := storage.NewSQLStore(db)

		business := biz.NewCreateEmployeeHandleOrderBiz(store)

		if err := business.CreateNewEmployeeHandleOrder(c.Request.Context(),&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}		
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}