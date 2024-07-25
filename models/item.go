package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	// ID          uint   `gorm:"primaryKey;autoIncrement"` // IDをプライマリーキーに設定
	Name        string `gorm:"not null"`
	Price       uint   `gorm:"not null"`
	Description string
	Soldout     bool `gorm:"not null;default:false"`
}
