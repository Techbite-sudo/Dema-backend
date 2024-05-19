package models

import (
	"time"
)

type User struct {
	Base
    FullName            string               `gorm:"not null"`
    Email               string               `gorm:"not null;unique"`
    AvatarUrl           string               `gorm:"not null;unique"`
    Gender              string               `gorm:"not null"`
    Weight              int              `gorm:"not null"`
    LastOnline          time.Time            `gorm:"not null"`
    DOB                 time.Time            `gorm:"not null"`
    Location            string               `gorm:"not null"`
    Password            string               `gorm:"not null"`
    Role                Role                 `gorm:"not null"`
    DietaryPreferences  []DietaryPreference  `gorm:"many2many:user_dietary_preferences;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    Allergies           []Allergy            `gorm:"many2many:user_allergies;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    Height              int
    ActivityLevel       string
    Favorites           []Favorite           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Role string

const (
    RoleUser  Role = "USER"
    RoleAdmin Role = "ADMIN"
)

type DietaryPreference struct {
    Base
    Name string `gorm:"not null"`
}

type Allergy struct {
    Base
    Name string `gorm:"not null"`
}

