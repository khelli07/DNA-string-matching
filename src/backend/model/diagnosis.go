package model

type Diagnosis struct {
	ID          int32   `gorm:"primaryKey"`
	InputDate   string  `json:"date"`
	Name        string  `json:"name" binding:"required"`
	DNASequence string  `json:"sequence" binding:"required"`
	DiseaseName string  `json:"disease" binding:"required"`
	Percentage  float32 `json:"percentage"`
	Result      bool    `json:"result"`
}
