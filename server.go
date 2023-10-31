package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/manuelxantony/blog/app"
	"github.com/manuelxantony/blog/controller"
	"github.com/manuelxantony/blog/database"
	"github.com/manuelxantony/blog/service"
)

var (
	appData        *app.App
	blogService    service.BlogService
	blogController controller.BlogController
)

func init() {
	config := LoadConfig()
	db, err := database.OpenDB(config)
	if err != nil {
		panic(err)
	}

	appData = &app.App{
		Config: config,
		DB:     db,
	}

	blogService = service.New(appData)
	blogController = controller.New(blogService)

}

func main() {

	server := gin.Default()

	blogApi := server.Group("/blog")
	{
		blogApi.GET("/showall", blogController.ShowAll)
		blogApi.GET("/showbyid", blogController.ShowById)
		blogApi.POST("/create", blogController.Create)
		blogApi.PUT("/update", blogController.Update)
		blogApi.DELETE("/delete", blogController.Delete)
	}

	port := appData.Config.Server.Port
	if port == "" {
		port = "8000"
	}
	server.Run(":" + port)
}
