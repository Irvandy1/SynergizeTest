package models

import "time"

type (
	LoginForm struct {
		UserID   int    `json:"user_id" gorm:"primaryKey"`
		Username string `json:"user_name"`
		Password string `json:"password"`
	}

	LogoutForm struct {
		UserID int `json:"user_id" gorm:"primaryKey"`
	}

	TopUpForm struct {
		UserID int     `json:"user_id"`
		Amount float64 `json:"amount"`
	}

	Session struct {
		UserID int       `json:"user_id"`
		Expiry time.Time `json:"expiry"`
	}

	User struct {
		UserID      int       `json:"user_id" gorm:"primaryKey"`
		UserName    string    `json:"username"`
		Password    string    `json:"password"`
		Email       string    `json:"email"`
		PhoneNumber string    `json:"phone_number"`
		RegisterAt  time.Time `json:"register_at"`
	}
	Bank struct {
		ID            int       `json:"id" gorm:"primaryKey"`
		BankCode      string    `json:"bank_code"`
		AccountNumber int       `json:"account_number"`
		AccountName   string    `json:"account_name"`
		UserID        int       `json:"user_id"`
		UpdatedAt     time.Time `json:"updated_at"`
	}

	Wallet struct {
		WalletID  int       `json:"wallet_id" gorm:"primaryKey"`
		Saldo     float64   `json:"saldo"`
		UpdatedAt time.Time `json:"updated_at"`
		UserID    int       `json:"user_id"`
	}

	UserDetail struct {
		User   User
		Bank   Bank
		Wallet Wallet
	}

	Pagination struct {
		Page  int    `json:"page"`
		Limit int    `json:"limit"`
		Keys  string `json:"keys"`
		Value string `json:"value"`
	}
)
