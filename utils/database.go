package utils

import (
	"Dema-backend/models"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitialiseDB() {
	var err error
	log.Print("Initialising Database...")

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Could not find Database Url")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	log.Print("Successfully connected database!")

	setupModels(
		&models.User{},
		// &models.Eatery{},
		// &models.Calories{},
		// &models.Allergy{},
		// &models.Compound{},
		// &models.DietaryPreference{},
		// &models.Favorite{},
		// &models.Food{},
		// &models.FoodName{},
		// &models.Ingredient{},
		// &models.Meal{},
		// &models.Nutrient{},
		// &models.RecipeStep{},
		// &models.Customer{},





		
	)
}

func setupModels(models ...interface{}) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		panic(err)
	}
}
