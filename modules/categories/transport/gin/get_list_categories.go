package ginCategories

import (
	"first-app/common"
	"first-app/modules/categories/biz"
	"first-app/modules/categories/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListCategories(db *gorm.DB) func(*gin.Context) {
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

		business := biz.NewGetListCategoriesBiz(store)
		result, err := business.GetNewListCategories(c.Request.Context(),&paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
