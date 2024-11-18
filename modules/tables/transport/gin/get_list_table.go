package ginTable

import (
	"first-app/common"
	"first-app/modules/tables/biz"
	"first-app/modules/tables/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListTable(db *gorm.DB) func(*gin.Context) {
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

		business := biz.NewGetListTableBiz(store)

		result, err := business.GetNewListTable(c.Request.Context(), &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
