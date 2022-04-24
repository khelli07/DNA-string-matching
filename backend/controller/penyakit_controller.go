package controller

import (
	"fmt"
	"net/http"

	mod "github.com/DNA-string-matching/backend/model"

	"github.com/gin-gonic/gin"
)

func AllPenyakitHandler(c *gin.Context) {
	var diseases []mod.Disease
	db := ConnectToDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured! Can't connect to DB"})
		return
	}

	res := db.Find(&diseases)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Penyakit tidak ditemukan!"})
	} else {
		response := gin.H{}
		for _, disease := range diseases {
			body := gin.H{
				"name":    disease.Name,
				"pattern": disease.Pattern,
			}

			response[fmt.Sprint(disease.ID)] = body
		}
		c.JSON(http.StatusOK, response)
	}

	return
}

func ReadPenyakitHandler(c *gin.Context) {
	var disease mod.Disease
	db := ConnectToDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured! Can't connect to DB"})
		return
	}

	res := db.Where("ID = ?", c.Param("id")).Find(&disease)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"name":    disease.Name,
			"pattern": disease.Pattern,
		})
	}

	return
}

func NewPenyakitHandler(c *gin.Context) {
	var disease mod.Disease
	db := ConnectToDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured! Can't connect to DB"})
		return
	}

	err := c.ShouldBindJSON(&disease)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint(err)})
	} else if !PatternIsValid(disease.Pattern) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pola tidak valid!"})
	} else {
		res := db.Create(&disease)
		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Model created successfully"})
		}
	}

	return
}

func UpdatePenyakitHandler(c *gin.Context) {
	var updatedDisease mod.Disease
	var disease mod.Disease
	db := ConnectToDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured! Can't connect to DB"})
		return
	}

	err := c.ShouldBindJSON(&updatedDisease)
	res := db.Where("ID = ?", c.Param("id")).Find(&disease)
	find := db.Model(&disease).Update("name", updatedDisease.Name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint(err)})
	} else if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
	} else if find.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": find.Error})
	} else if !PatternIsValid(updatedDisease.Pattern) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pola tidak valid!"})
	} else {
		db.Model(&disease).Update("pattern", updatedDisease.Pattern)
		c.JSON(http.StatusOK, gin.H{"message": "Model updated successfully"})
	}

	return
}

func DeletePenyakitHandler(c *gin.Context) {
	var disease mod.Disease
	db := ConnectToDB()
	db.Model(&disease).Where("ID = ?", c.Param("id")).Delete(&disease)

	return
}
