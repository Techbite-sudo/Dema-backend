package models

type Eatery struct {
    Base
    Name           string    `gorm:"not null"`
    Location       string    `gorm:"not null"`
    Cuisine        string    `gorm:"not null"`
    ContactDetails string    `gorm:"not null"`
    MenuItems      []Meal    `gorm:"many2many:eatery_menu_items;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}