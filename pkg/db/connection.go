package db

import (
	config "ecommerce/pkg/config"
	domain "ecommerce/pkg/domain"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBPort, cfg.DBName)

	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.Inventories{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Cart{})
	db.AutoMigrate(&domain.Wishlist{})
	db.AutoMigrate(&domain.WishlistItems{})
	db.AutoMigrate(&domain.Address{})
	db.AutoMigrate(&domain.Order{})
	db.AutoMigrate(&domain.OrderItem{})
	db.AutoMigrate(&domain.LineItems{})
	db.AutoMigrate(&domain.PaymentMethod{})
	db.AutoMigrate(&domain.Offer{})
	db.AutoMigrate(&domain.Coupons{})
	db.AutoMigrate(&domain.Wallet{})
	db.AutoMigrate(&domain.WalletHistory{})
	db.AutoMigrate(&domain.Image{})

	return db, dbErr
}
