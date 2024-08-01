package main

import (
	"gin-web-app/controllers"
	"gin-web-app/infra"
	"gin-web-app/middlewares"
	"gin-web-app/repositories"
	"gin-web-app/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	r.Use(cors.Default())
	itemRouter := r.Group("/items")
	itemRouterWithAuth := r.Group("/items", middlewares.AuthMiddleware(authService))
	authRouter := r.Group("/auth")

	itemRouter.GET("", itemController.FindAll)
	itemRouterWithAuth.GET("/:id", itemController.FindById)  // パスパラメータを受け取る
	itemRouterWithAuth.POST("", itemController.Create)       // POSTリクエストを受け取り、アイテムの追加を行う
	itemRouterWithAuth.PUT("/:id", itemController.Update)    // PUTリクエストを受け取り、アイテムの更新を行う
	itemRouterWithAuth.DELETE("/:id", itemController.Delete) // DELETEリクエストを受け取り、アイテムの削除を行う

	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	return r
}

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	r := setupRouter(db)

	r.Run("localhost:8080")
}
