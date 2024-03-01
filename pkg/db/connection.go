package db

import (
	"ecommerce/pkg/config"
	domain "ecommerce/pkg/domain"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	// Construct the connection string using the configuration
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)

	// Open a connection to the database
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})
	if dbErr != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", dbErr)
	}

	// Auto migrate the database tables
	if err := db.AutoMigrate(
		&domain.Inventories{},
		&domain.Category{},
		&domain.Users{},
		&domain.Admin{},
		domain.Cart{},
		domain.Address{},
		domain.Order{},
		domain.OrderItem{},
		domain.PaymentMethod{},
		domain.Coupons{},
		domain.Wallet{},
		domain.Offer{},
		domain.LineItems{},
		domain.Wishlist{},
	); err != nil {
		return db, fmt.Errorf("failed to auto-migrate database tables: %v", err)
	}

	// Check if admin user exists and create one if not
	if err := CheckAndCreateAdmin(db); err != nil {
		return db, fmt.Errorf("failed to check and create admin user: %v", err)
	}

	return db, nil
}

func CheckAndCreateAdmin(db *gorm.DB) error {
	var count int64
	db.Model(&domain.Admin{}).Count(&count)

	fmt.Println("error occured here")
	
	if count == 0 {
		password := "password123"
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			fmt.Println("check and create admin error")
			return err
		}
		admin := domain.Admin{
			ID:       1,
			Name:     "admin",
			UserName: "animestore@gmal.com",
			Password: string(hashedPass),
		}
		db.Create(&admin)
	}
	return nil
}
