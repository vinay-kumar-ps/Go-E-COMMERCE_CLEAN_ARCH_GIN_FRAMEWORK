package repository

// import (
// 	"ecommerce/pkg/domain"
// 	"ecommerce/pkg/repository/interfaces"
// 	"errors"
// 	"time"

// 	"gorm.io/gorm"
// )

// type walletRepository struct {
// 	DB *gorm.DB
// }

// //
// func NewWalletRepository(DB *gorm.DB) interfaces.WalletRepository {
// 	return &walletRepository{
// 		DB: DB,
// 	}
// }
// func (wr *walletRepository) CreditToUserWallet(amount float64, walletId int) error {
// 	if err := wr.DB.Exec("UPDSTE wallets SET amount=amount+$1 WHERE wallet_id=$2", amount, walletId).Error; err != nil {
// 		return errors.New("amount adding to wallet failed")
// 	}
// 	return nil
// }

// func (wr *walletRepository) FindUserIdFromOrderId(id int) (int, error) {
// 	var userId int
// 	err := wr.DB.Raw("SELECT user_id FROM orders WHERE order_id=?", id).Scan(&userId).Error
// 	if err != nil {
// 		return 0, errors.New("user id not found")
// 	}
// 	return userId, nil
// }

// func (wr *walletRepository) FindWalletIdFromUserId(userId int) (int, error) {
// 	var walletCount int
// 	if err := wr.DB.Raw("SELECT COUNT (*)FROM wallets WHERE user_id=?", userId).Scan(&walletCount).Error; err != nil {
// 		return 0, errors.New("wallet not found")
// 	}
// 	var walletId int
// 	if walletCount > 0 {
// 		err := wr.DB.Raw("SELECT if FROM wallets WHERE user_id =?", userId).Scan(&walletId).Error
// 		if err != nil {
// 			return 0, errors.New("wallet id not found")
// 		}

// 	}
// 	return walletId, nil
// }
// func (wr *walletRepository) CreateNewWallet(userId int) (int, error) {
// 	var walletId int
// 	err := wr.DB.Exec("INSERT INTO wallets (user_id,amount)VALUES (?,?)", userId, 0).Error
// 	if err != nil {
// 		return 0, err
// 	}
// 	if err := wr.DB.Raw("SELECT id FROM wallet WHERE user_id=$1", userId).Scan(&walletId).Error; err != nil {
// 		return 0, err
// 	}
// 	return walletId, nil
// }
// func (wr *walletRepository) GetBalance(walletId int) (int, error) {
// 	var balance float64

// 	if err := wr.DB.Raw("SELECT amount FROM wallets WHERE id=?", walletId).Scan(&balance).Error; err != nil {
// 		return 0, errors.New("balance not found")

// 	}
// 	return int(balance), nil
// }
// func (wr *walletRepository) GetHistory(walletId, page, limit int) ([]domain.WalletHistory, error) {
// 	var walletHistory []domain.WalletHistory
// 	if page == 0 {
// 		page = 1
// 	}
// 	if limit == 0 {
// 		limit = 10
// 	}
// 	offset := (page - 1*limit)
// 	err := wr.DB.Raw("SELECT * FROM wallet_histories WHERE wallet_id=? limit ? offset ?", walletId, limit, offset).Scan(&walletHistory).Error

// 	if err != nil {
// 		return []domain.WalletHistory{}, errors.New("wallet history not found")

// 	}
// 	return walletHistory, nil
// }
// func (wr *walletRepository) AddHistory(amount, walletId int, purpose string) error {
// 	err := wr.DB.Exec("INSERT INTO wallet_histories (wallet_id,amount,purpose,time)VALUES(?,?,?,?)", walletId, amount, purpose, time.Now()).Error

// 	if err != nil {
// 		return errors.New("history adding  failed")
// 	}

// 	return nil
// }
// func (wr *walletRepository) PayFromWallet(userId, orderId int, price float64) (float64, error) {
// 	var balanceAfterPay float64
// 	if err := wr.DB.Exec("UPDATE wallets SET amount =? WHERE user_id=?", price, userId).Error; err != nil {

// 		return 0, err
// 	}
// 	if err := wr.DB.Exec("UPDATE orders SET payment_status='PAID' WHERE order_id=?", orderId).Error; err != nil {
// 		return 0, err
// 	}
// 	if err := wr.DB.Raw("SELECT amount FROM wallets WHERE user_id=?", userId).Scan(&balanceAfterPay).Error; err != nil {
// 		return 0, err
// 	}
// 	return balanceAfterPay, nil

// }
