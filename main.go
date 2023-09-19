package main

import (
	"fmt"

	"github.com/Bobs-code/Guitar-API/controllers"
	"github.com/gin-gonic/gin"
)

func getGuitars(c *gin.Context) {
	guitars := controllers.GetGuitars()
	c.JSON(200, guitars)
}

func homePage(c *gin.Context) {
	c.Writer.WriteString("GuitarAPI Project Home Page")
	fmt.Println("Endpoint Hit: Home Page")
}

func main() {

	r := gin.Default()
	r.GET("/", homePage)
	r.GET("/guitars", getGuitars)

	r.Run("localhost:8080")

}
