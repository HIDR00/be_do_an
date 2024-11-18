package main

import (
	ginCategories "first-app/modules/categories/transport/gin"
	ginDevices "first-app/modules/devices/transport/gin"
	ginEmployeeHandleOrder "first-app/modules/employee_handle_order/transport/gin"
	ginMenu "first-app/modules/menus/transport/gin"
	ginOrder "first-app/modules/orders/transport/gin"
	ginTable "first-app/modules/tables/transport/gin"
	ginUser "first-app/modules/users/transport/gin"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/Go Begin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// CRUD
	// POST /v1/items (create a new item)
	// GET /v1/items (list items) /v1/items?page=1
	// GET /v1/items/:id (get item detail by id)
	// PUSH v1/items/:id (update an item by id)
	// DELETE /v1/items/:id (delete item by id)

	v1 := r.Group("v1")
	{
		menus := v1.Group("/menus")
		{
			menus.POST("", ginMenu.CreateMenu(db))
			menus.GET("", ginMenu.ListItem(db))
			menus.GET(`/:id`, ginMenu.GetMenuById(db))
			menus.PATCH("/:id", ginMenu.UpdateMenu(db))
			menus.DELETE("/:id", ginMenu.DeleteMenu(db))
		}
		tables := v1.Group("/tables")
		{
			tables.POST("", ginTable.CreateTable(db))
			tables.GET("", ginTable.ListTable(db))
			tables.GET(`/:id`, ginTable.GetTableById(db))
			tables.PATCH("/:id", ginTable.UpdateTable(db))
			// images.DELETE("/:id", ginItem.DeleteItem(db))
		}
		orders := v1.Group("/orders")
		{
			orders.POST("/:id", ginOrder.CreateOrder(db))
			orders.GET("/total", ginOrder.GetTotalOrder(db))
			orders.GET(`/:id`, ginOrder.GetOrderById(db))
			// orders.PATCH("/:id", ginOrder.UpdateOrder(db))
			orders.DELETE("/:id", ginOrder.DeleteOrder(db))
		}
		categories := v1.Group("/categories")
		{
			// categories.POST("/:id", ginOrder.CreateOrder(db))
			categories.GET("", ginCategories.ListCategories(db))
			// categories.GET(`/:id`, ginOrder.GetOrderById(db))
			// categories.PATCH("/:id", ginOrder.UpdateOrder(db))
			// images.DELETE("/:id", ginItem.DeleteItem(db))
		}
		users := v1.Group("/users")
		{
			users.POST("/auth", ginUser.AuthenUser(db))
			users.POST("", ginUser.CreateUser(db))
			users.GET("", ginUser.ListUser(db))
			users.GET(`/:id`, ginUser.GetUserById(db))
			users.PATCH("/:id", ginUser.UpdateUser(db))
			users.DELETE("/:id", ginUser.DeleteUser(db))
		}
		devices := v1.Group("/devices")
		{
			devices.GET("", ginDevices.ListDevices(db))
			devices.POST("", ginDevices.CreateDevices(db))
		}
		employeeHandleOrder := v1.Group("/employee_handle_order")
		{
			employeeHandleOrder.POST("", ginEmployeeHandleOrder.CreateEmployeeHandleOrder(db))
			employeeHandleOrder.GET("", ginEmployeeHandleOrder.ListEmployeeHandleOrder(db))
		}
	}
	r.Run()
}
