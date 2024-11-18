package ginMenu


import (
	"first-app/common"
	"first-app/modules/menus/biz"
	"first-app/modules/menus/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListItem(db *gorm.DB) func(*gin.Context) {
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

		business := biz.NewGetListMenuBiz(store)

		result, err := business.GetNewListMenu(c.Request.Context(), &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
