package models

var (
	User struct {
		UserID      string  `json:"user_id"`
		Email       string  `json:"email"`
		PhoneNumber string  `json:"phone_number"`
		Saldo       float64 `json:"saldo"`
		RegisterAt  string  `json:"register_at"`
	}

	Bank struct {
		BankName          string `json:"bank_name"`
		BankAccountNumber string `json:"bank_account_number"`
		BankAccountName   string `json:"bank_account_name"`
	}
)
