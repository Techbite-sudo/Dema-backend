package models

import(
	"time"
	uuid "github.com/satori/go.uuid"

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
    Base
    MealID   uuid.UUID `gorm:"type:uuid"`
    Step        int    `gorm:"not null"`
    Description string `gorm:"not null"`
}

type Ingredient struct {
    Base
    MealID   uuid.UUID `gorm:"type:uuid"`
    Name   string `gorm:"not null"`
    Amount string `gorm:"not null"`
}