package main

import (
	"example.com/m/controllers"
	"example.com/m/initializers"
	model "example.com/m/models"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.DB.AutoMigrate(&model.Cat{})
}

func main()  {
	r := gin.Default()

	r.POST("/", controllers.CatCreate)
	r.GET("/", controllers.FindAll)
	r.GET("/:id", controllers.FindOne)
	r.PUT("/:id", controllers.UpdateOne)
	r.DELETE("/:id", controllers.DeleteOne)
	
	r.Run()
}