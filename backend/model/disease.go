package model

type Disease struct {
	ID      int32  `gorm:"primaryKey"`
	Name    string `json:"name" binding:"required" gorm:"unique"`
	Pattern string `json:"pattern" binding:"required"`
}
