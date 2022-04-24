package controller

import (
	"fmt"
	"net/http"
	"time"

	mod "github.com/DNA-string-matching/backend/model"

	"github.com/gin-gonic/gin"
)

func AllDiagnosisHandler(c *gin.Context) {
	var diagnoses []mod.Diagnosis
	db := ConnectToDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured! Can't connect to DB"})
		return
	}

	query := fmt.Sprint(c.Query("query"))
	date, name := ExtractQuery(query)

	res := db.Find(&diagnoses)
	if date != "" || name != "" {
		name_pattern := "%" + name + "%"
		res = db.Where("input_date = ? OR disease_name = ? OR name like ?",
			date, name, name_pattern).Find(&diagnoses)
	}

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Penyakit tidak ditemukan!"})
	} else {
		response := gin.H{}
		for _, diagnose := range diagnoses {
			body := gin.H{
				"date":       diagnose.InputDate,
				"name":       diagnose.Name,
				"sequence":   diagnose.DNASequence,
				"disease":    diagnose.DiseaseName,
				"percentage": diagnose.Percentage,
				"result":     diagnose.Result,
			}

			response[fmt.Sprint(diagnose.ID)] = body
		}
		c.JSON(http.StatusOK, response)
	}
}

func ReadDiagnosisHandler(c *gin.Context) {
	var diagnosis mod.Diagnosis
	db := ConnectToDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured! Can't connect to DB"})
		return
	}

	res := db.Where("ID = ?", c.Param("id")).Find(&diagnosis)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Penyakit tidak ditemukan!"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"date":       diagnosis.InputDate,
			"name":       diagnosis.Name,
			"sequence":   diagnosis.DNASequence,
			"disease":    diagnosis.DiseaseName,
			"percentage": diagnosis.Percentage,
			"result":     diagnosis.Result,
		})
	}
	return
}

func NewDiagnosisHandler(c *gin.Context) {
	var diagnose mod.Diagnosis
	db := ConnectToDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured! Can't connect to DB"})
		return
	}

	err := c.ShouldBindJSON(&diagnose)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint(err)})
	} else if !PatternIsValid(diagnose.DNASequence) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pola tidak valid!"})
	} else {
		diagnose.InputDate = time.Now().Format("2006-01-02")
		//TODO: CHANGE TO FUNCTION, NOT HARDCODED
		diagnose.Percentage = 0.8
		diagnose.Result = true

		res := db.Create(&diagnose)
		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"new_id":  diagnose.ID,
				"message": "Model created successfully",
			})
		}
	}

	return
}

func DeleteDiagnosisHandler(c *gin.Context) {
	var diagnosis mod.Diagnosis
	db := ConnectToDB()
	db.Model(&diagnosis).Where("ID = ?", c.Param("id")).Delete(&diagnosis)
	return
}
