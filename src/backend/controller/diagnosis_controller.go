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

	res := db.Find(&diagnoses)
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

func SearchDiagnosisHandler(c *gin.Context) {
	type Query struct {
		Query string `json:"query"`
	}
	var query Query
	var diagnoses []mod.Diagnosis
	db := ConnectToDB()

	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured! Can't connect to DB"})
		return
	}

	err := c.ShouldBindJSON(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint(err)})
	} else {
		date, name := ExtractQuery(query.Query)

		res := db.Find(&diagnoses)
		if date != "" || name != "" {
			if date != "" && name != "" {
				res = db.Where("input_date = ? AND disease_name = ?", date, name).Find(&diagnoses)
			} else if date != "" {
				res = db.Where("input_date = ?", date).Find(&diagnoses)
			} else {
				res = db.Where("disease_name = ?", name).Find(&diagnoses)

			}
		}

		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint(res.Error)})
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
	type DiagnoseInput struct {
		Name         string `json:"name" binding:"required"`
		DNASequence  string `json:"sequence" binding:"required"`
		DiseaseName  string `json:"disease" binding:"required"`
		AlgorithmIdx int    `json:"algo_index" binding:"required"`
	}

	var userInput DiagnoseInput
	db := ConnectToDB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured! Can't connect to DB"})
		return
	}

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint(err)})
	} else if !PatternIsValid(userInput.DNASequence) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pola tidak valid!"})
	} else {
		var diagnose mod.Diagnosis
		var disease mod.Disease
		var exists bool

		db.Model(&disease).Select("count(*) > 0").Where("name = ?", userInput.DiseaseName).Find(&exists)
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Disease not found!"})
		} else {
			db.First(&disease, "name = ?", userInput.DiseaseName)
			diagnose.Name = userInput.Name
			diagnose.DNASequence = userInput.DNASequence
			diagnose.DiseaseName = userInput.DiseaseName
			diagnose.InputDate = time.Now().Format("2006-01-02")
			diagnose.Percentage, diagnose.Result = DNAStringMatching(disease.Pattern, userInput.DNASequence, userInput.AlgorithmIdx-1)
			res := db.Create(&diagnose)

			if res.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"id":      diagnose.ID,
					"message": "Model created successfully"})
			}
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
