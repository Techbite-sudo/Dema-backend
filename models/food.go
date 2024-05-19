package models
import(
	"time"
)
type Food struct {
    Base
    ImageUrl      string         `gorm:"not null"`
    CommonName    string         `gorm:"not null"`
    Description   string         `gorm:"not null"`
    FoodGroup     string         `gorm:"not null"`
    LastUpdate    time.Time      `gorm:"not null"`
    Calories      Calories       `gorm:"not null"`
    OtherNames    []FoodName     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    Sources       []Source       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    Compounds     []Compound     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    Nutrients     []Nutrient     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    RegionsFound  []string
    Category      string
    ServingSize   string
    Verified      bool           `gorm:"not null"`
}

type Calories struct {
    Amount int `gorm:"not null"`
}

type FoodName struct {
    Category string `gorm:"not null"`
    Name     string `gorm:"not null"`
}

type Source struct {
    Name    string `gorm:"not null"`
    UrlLink string `gorm:"not null"`
}

type Compound struct {
    Name   string  `gorm:"not null"`
    Amount float64 `gorm:"not null"`
}

type Nutrient struct {
    Name   string  `gorm:"not null"`
    Amount float64 `gorm:"not null"`
}
