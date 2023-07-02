package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"network.com/network/config"
	"network.com/network/models"
)

var db *gorm.DB

func main() {

	println("Runnning network service")
	startDbJobs()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		var data models.Company
		db.Find(&data)
		c.JSON(http.StatusOK, data)
	})

	r.Run(":8080")

}

func startDbJobs() {
	config.InitialiseDbConnection()
	db = config.GetDbConnection()
	models.InitCompanyMigration()
}
