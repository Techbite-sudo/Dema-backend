package models

import(
	"time"
)
type Customer struct {
    Base
    DOB           time.Time  `gorm:"not null"`
    Weight        int        `gorm:"not null"`
    Height        int        `gorm:"not null"`
    FoodPreference string     `gorm:"not null"`
    Allergies     []string
    NumberOfMeals int        `gorm:"not null"`
    Meals         []string   `gorm:"not null"`
    History       []string   `gorm:"not null"`
}