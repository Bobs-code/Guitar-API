package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Bobs-code/Guitar-API/controllers"
	"github.com/Bobs-code/Guitar-API/models"
	"github.com/gin-gonic/gin"
)

func getGuitars(c *gin.Context) {
	guitars := controllers.GetGuitars()
	c.JSON(200, guitars)
}
func singleGuitar(c *gin.Context) {

	urlId := c.Param("id")

	urlIdInt, err := strconv.Atoi(urlId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	guitar, err := controllers.GetSingleGuitar(urlIdInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guitar not found"})
	}

	c.JSON(http.StatusOK, guitar)
}
func addGuitar(c *gin.Context) {
	var guitar models.Guitar
	if err := c.ShouldBindJSON(&guitar); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newGuitar := models.Guitar{
		Brand_id:    guitar.Brand_id,
		Model:       guitar.Model,
		Year:        guitar.Year,
		Description: guitar.Description,
	}
	db, err := controllers.InitPGDB()
	if err != nil {
		log.Println(err)
	}
	if err := db.Create(&newGuitar).Error; err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Database Migration Error"})
	}
	c.JSON(http.StatusOK, guitar)

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
	r.GET("/guitars/:id", singleGuitar)
	r.PUT("/guitars/add", addGuitar)

	// nolint:errcheck
	r.Run("localhost:8080")

}
