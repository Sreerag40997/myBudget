package seed

import (
	"log"

	"github.com/sreerag/myBudget/config"
	"github.com/sreerag/myBudget/models"
	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin() {
	var admin models.User
	if err := config.DB.Where("email = ?", "admin@gmail.com").First(&admin).Error; err == nil {
		log.Println("Admin already exists")
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin = models.User{
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Password: string(hashed),
		Role:     1,
	}
	config.DB.Create(&admin)
	log.Println("Admin created with email: admin@gmail.com, password: admin123")
}
