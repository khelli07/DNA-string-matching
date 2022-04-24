package main

import (
	"log"
	"net/http"

	ctrl "github.com/DNA-string-matching/backend/controller"
	mod "github.com/DNA-string-matching/backend/model"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":   "Stima",
		"school": "Institut Teknologi Bandung",
	})

}

func main() {
	dsn := "root:" + ctrl.Password + "@tcp(127.0.0.1:3306)/gosql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB conn error")
	}

	db.AutoMigrate(&mod.Disease{})
	db.AutoMigrate(&mod.Diagnosis{})

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", rootHandler)

	penyakit := router.Group("/penyakit")
	{
		penyakit.GET("", ctrl.AllPenyakitHandler)
		penyakit.GET("/:id", ctrl.ReadPenyakitHandler)
		penyakit.POST("/new", ctrl.NewPenyakitHandler)
		penyakit.PATCH("/:id", ctrl.UpdatePenyakitHandler)
		penyakit.DELETE("/:id/delete", ctrl.DeletePenyakitHandler)
	}

	diagnosis := router.Group("/diagnosis")
	{
		diagnosis.GET("", ctrl.AllDiagnosisHandler)
		diagnosis.GET("/:id", ctrl.ReadDiagnosisHandler)
		diagnosis.POST("/new", ctrl.NewDiagnosisHandler)
		diagnosis.DELETE("/:id/delete", ctrl.DeleteDiagnosisHandler)
	}

	router.Run(":8080")
}
