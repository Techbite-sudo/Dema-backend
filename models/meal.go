package models

import(
	"time"
)

type Meal struct {
    Base
    DocName        string         `gorm:"not null"`
    ImageUrl       string         `gorm:"not null"`
    CommonName     string         `gorm:"not null"`
    Description    string         `gorm:"not null"`
    LastUpdate     time.Time      `gorm:"not null"`
    Foods          []Food         `gorm:"many2many:meal_foods;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    OtherNames     []string       `gorm:"not null"`
    Recipe         []RecipeStep   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    Ingredients    []Ingredient   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    Cuisine        string
    PreparationTime string
}

type RecipeStep struct {
    Step        int    `gorm:"not null"`
    Description string `gorm:"not null"`
}

type Ingredient struct {
    Name   string `gorm:"not null"`
    Amount string `gorm:"not null"`
}