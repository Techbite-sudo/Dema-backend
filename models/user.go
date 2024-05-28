package models

import (
	"time"
)

type User struct {
    Base
	FullName    string      `gorm:"not null"`
	Email       string      `gorm:"not null;unique"`
	Password    string      `gorm:"not null"`
	AvatarUrl   string      `gorm:"not null;unique"`
	Gender      string      `gorm:"not null"`
	LastOnline  time.Time   `gorm:"not null"`
	Role        Role        `gorm:"not null"`
	AccountType AccountType `gorm:"not null"`
	Location    string      `gorm:"not null"`
}

type Role string

const (
	RoleAdmin     Role = "ADMIN"
	RoleDataEntry Role = "DATA_ENTRY"
	RoleCustomer  Role = "CUSTOMER"
	RoleEatery    Role = "EATERY"
)

type AccountType string

const (
	AccountTypeIndividual AccountType = "INDIVIDUAL"
	AccountTypeBusiness   AccountType = "BUSINESS"
)

type DietaryPreference struct {
	Base
	Name string `gorm:"not null"`
}

type Allergy struct {
	Base
	Name string `gorm:"not null"`
}

// package models

// import (
// 	"time"
// )

// type User struct {
// 	Base
//     FullName            string               `gorm:"not null"`
//     Email               string               `gorm:"not null;unique"`
//     AvatarUrl           string               `gorm:"not null;unique"`
//     Gender              string               `gorm:"not null"`
//     Weight              int              `gorm:"not null"`
//     LastOnline          time.Time            `gorm:"not null"`
//     DOB                 time.Time            `gorm:"not null"`
//     Location            string               `gorm:"not null"`
//     Password            string               `gorm:"not null"`
//     Role                Role                 `gorm:"not null"`
//     DietaryPreferences  []DietaryPreference  `gorm:"many2many:user_dietary_preferences;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
//     Allergies           []Allergy            `gorm:"many2many:user_allergies;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
//     Height              int
//     ActivityLevel       string
//     Favorites           []Favorite           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// }

// type Role string

// const (
//     RoleAdmin Role = "ADMIN"
//     RoleCustomer  Role = "CUSTOMER"
//     RoleDataEntry  Role = "DATAENTRY"
//     RoleEatery Role = "EATERY"
// )
