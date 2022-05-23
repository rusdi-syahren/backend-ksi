package external

// PaymentResponse struct
type PaymentResponse struct {
	StatusCode        string     `json:"status_code"`
	StatusMessage     string     `json:"status_message"`
	TransactionID     string     `json:"transaction_id"`
	OrderID           string     `json:"order_id"`
	MerchantID        string     `json:"merchant_id"`
	GrossAmount       string     `json:"gross_amount"`
	Currency          string     `json:"currency"`
	PaymentType       string     `json:"payment_type"`
	TransactionTime   string     `json:"transaction_time"`
	TransactionStatus string     `json:"transaction_status"`
	FraudStatus       string     `json:"fraud_status"`
	Action            ActionList `json:"actions"`
	PermataVANumber   string     `json:"permata_va_number"`
}

// Action struct
type Action struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	URL    string `json:"url"`
}

// ActionList struct
type ActionList []Action

//PaymentCharge struct
type PaymentCharge struct {
	PaymentType          string             `json:"payment_type"`
	TransacactionDetails TransactionDetails `json:"transaction_details"`
	ItemDetails          ItemDetails        `json:"item_details"`
	CustomerDetails      CustomerDetails    `json:"customer_details"`
}

//BankTransfer struct
type BankTransfer struct {
	Bank string `json:"bank"`
}

//TransactionDetails struct
type TransactionDetails struct {
	OrderID     string  `json:"order_id"`
	GrossAmount float64 `json:"gross_amount"`
}

//ItemDetail struct
type ItemDetail struct {
	ID       string  `json:"id"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Name     string  `json:"name"`
}

// ItemDetails struct
type ItemDetails []*ItemDetail

//CustomerDetails struct
type CustomerDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

// PaymentChargeGopay struct
type PaymentChargeGopay struct {
	PaymentCharge
	Gopay ChargeGopay `json:"gopay"`
}

// PaymentChargeCC struct
type PaymentChargeCC struct {
	PaymentCharge
	ChargeCC ChargeCC `json:"credit_card"`
}

// PaymentChargePermata struct
type PaymentChargePermata struct {
	PaymentCharge
	BankTransfer BankTransfer `json:"bank_transfer"`
}

//ChargeGopay struct
type ChargeGopay struct {
	EnableCallback string `json:"enable_callback"`
	CallbackURL    string `json:"callback_url"`
}

// ChargeCC struct
type ChargeCC struct {
	TokenID string `json:"token_id"`
}

// TokenParams struct
type TokenParams struct {
	GrossAmount  string `json:"gross_amount"`
	CardNumber   string `json:"card_number"`
	CardExpMonth string `json:"card_exp_month"`
	CardExpYear  string `json:"card_exp_year"`
	CardCVV      string `json:"card_cvv"`
}

// TokenResponse struct
type TokenResponse struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_message"`
	TokenID       string `json:"token_id"`
	Hash          string `json:"hash"`
}
