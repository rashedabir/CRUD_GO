package controllers

import (
	"example.com/m/initializers"
	model "example.com/m/models"
	"github.com/gin-gonic/gin"
)

func CatCreate(c *gin.Context){
	var body struct{
		Name string
	}
	c.Bind((&body))

	cat := model.Cat{Name: body.Name}
	result := initializers.DB.Create(&cat)

	if(result.Error != nil){
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"data": cat,
	})
}

func FindAll(c *gin.Context){
	var cats []model.Cat
	result := initializers.DB.Find(&cats)

	if(result.Error != nil){
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"data": cats,
	})
}

func FindOne(c *gin.Context){
	id := c.Param("id")
	
	var cat model.Cat
	result := initializers.DB.First(&cat, id)

	if(result.Error != nil){
		c.JSON(400, gin.H{
            "error": "Cat not found",
        })
		return
	}

	c.JSON(200, gin.H{
		"data": cat,
	})
}

func UpdateOne(c *gin.Context){
	id := c.Param("id")
	var body struct{
		Name string
	}
	c.Bind((&body))
	
	var cat model.Cat
	result := initializers.DB.First(&cat, id)

	if(result.Error != nil){
		c.JSON(400, gin.H{
            "error": "Cat not found",
        })
		return
	}

	initializers.DB.Model(&cat).Updates(model.Cat{
		Name: body.Name,
	})

	c.JSON(200, gin.H{
		"data": cat,
	})
}


func DeleteOne(c *gin.Context){
	id := c.Param("id")
	
	var cat model.Cat
	result := initializers.DB.First(&cat, id)

	if(result.Error != nil){
		c.JSON(400, gin.H{
            "error": "Cat not found",
        })
		return
	}

	initializers.DB.Delete(&model.Cat{}, id)

	c.Status(200)
}