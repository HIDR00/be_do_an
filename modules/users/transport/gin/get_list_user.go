package ginUser 

import (
	"first-app/common"
	"first-app/modules/users/biz"
	"first-app/modules/users/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListUser(db *gorm.DB) func(*gin.Context) {
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

		business := biz.NewGetListUserBiz(store)
		result, err := business.GetNewListUser(c.Request.Context(), &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
