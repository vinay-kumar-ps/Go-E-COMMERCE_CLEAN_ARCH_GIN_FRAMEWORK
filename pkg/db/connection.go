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
        return nil, dbErr
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
        return db, err
    }

    // Check if admin user exists and create one if not
    CheckAndCreateAdmin(db)

    return db, nil
}

func CheckAndCreateAdmin(db *gorm.DB) {
    var count int64
    db.Model(&domain.Admin{}).Count(&count)
    if count == 0 {
        // If no admin user exists, create one with default credentials
        password := "123"
        hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
        if err != nil {
            return
        }

        admin := domain.Admin{
            ID:       1,
            Username: "admin", // Change this to your desired username
            Password: string(hashedPassword),
        }
        db.Create(&admin)
    }
}
