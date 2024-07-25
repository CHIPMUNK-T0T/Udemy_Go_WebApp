package main

import (
	"gin-web-app/controllers"
	"gin-web-app/infra"
	"gin-web-app/repositories"
	"gin-web-app/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	// items := []models.Item{
	// 	{ID: 1, Name: "商品1", Price: 1000, Description: "説明１", Soldout: false},
	// 	{ID: 2, Name: "商品2", Price: 2000, Description: "説明２", Soldout: true},
	// 	{ID: 3, Name: "商品3", Price: 3000, Description: "説明３", Soldout: false},
	// }

	// itemRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()

	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)  // パスパラメータを受け取る
	r.POST("/items", itemController.Create)       // POSTリクエストを受け取り、アイテムの追加を行う
	r.PUT("/items/:id", itemController.Update)    // PUTリクエストを受け取り、アイテムの更新を行う
	r.DELETE("/items/:id", itemController.Delete) // DELETEリクエストを受け取り、アイテムの削除を行う

	r.Run("localhost:8080")
}
