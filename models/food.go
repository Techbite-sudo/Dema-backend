package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Food struct {
	Base
	ImageUrl     string     `gorm:"not null"`
	CommonName   string     `gorm:"not null"`
	Description  string     `gorm:"not null"`
	FoodGroup    string     `gorm:"not null"`
	LastUpdate   time.Time  `gorm:"not null"`
	Calories     Calories   `gorm:"foreignKey:FoodID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OtherNames   []FoodName `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:FoodID;"`
	Sources      []Source   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:FoodID;"`
	Compounds    []Compound `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:FoodID;"`
	Nutrients    []Nutrient `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:FoodID;"`
	RegionsFound []string
	Category     string
	ServingSize  string
	Verified     bool `gorm:"not null"`
}

type Calories struct {
	Base
	FoodID   uuid.UUID `gorm:"type:uuid"`
	Amount int `gorm:"not null"`
}

type FoodName struct {
	Base
	FoodID   uuid.UUID `gorm:"type:uuid"`
	Category string    `gorm:"not null"`
	Name     string    `gorm:"not null"`
}

type Source struct {
	Base
	FoodID  uuid.UUID `gorm:"type:uuid"`
	Name    string    `gorm:"not null"`
	UrlLink string    `gorm:"not null"`
}

type Compound struct {
	Base
	FoodID uuid.UUID `gorm:"type:uuid"`
	Name   string    `gorm:"not null"`
	Amount float64   `gorm:"not null"`
}

type Nutrient struct {
	Base
	FoodID uuid.UUID `gorm:"type:uuid"`
	Name   string    `gorm:"not null"`
	Amount float64   `gorm:"not null"`
}
