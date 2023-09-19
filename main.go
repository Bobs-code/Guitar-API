package main

import (
	"fmt"
	"net/http"

	"github.com/Bobs-code/Guitar-API/controllers"
	"github.com/gin-gonic/gin"
)

func getGuitars(c *gin.Context) {
	guitars := controllers.GetGuitars()
	c.JSON(200, guitars)
}

func homePage(c *gin.Context) {
	_, err := c.Writer.WriteString("GuitarAPI Project Home Page")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Endpoint Hit: Home Page")
}

func main() {

	r := gin.Default()
	r.GET("/", homePage)
	r.GET("/guitars", getGuitars)

	// nolint:errcheck
	r.Run("localhost:8080")

}
