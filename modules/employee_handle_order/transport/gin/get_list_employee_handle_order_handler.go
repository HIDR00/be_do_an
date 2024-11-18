package ginEmployeeHandleOrder 

import (
	"first-app/common"
	"first-app/modules/employee_handle_order/biz"
	"first-app/modules/employee_handle_order/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListEmployeeHandleOrder(db *gorm.DB) func(*gin.Context) {
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

		business := biz.NewGetListEmployeeHandleOrderBiz(store)
		result, err := business.GetNewListEmployeeHandleOrder(c.Request.Context(), &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
