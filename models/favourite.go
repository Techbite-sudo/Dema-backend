package models
import(
	"time"
	uuid "github.com/satori/go.uuid"
)

type Favorite struct {
    Base
    UserID    uuid.UUID   `gorm:"type:uuid"`
    User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
    FoodID    uuid.UUID   `gorm:"type:uuid"`
    Food      Food      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    MealID    uuid.UUID   `gorm:"type:uuid"`
    Meal      Meal      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    Timestamp time.Time `gorm:"not null"`
}