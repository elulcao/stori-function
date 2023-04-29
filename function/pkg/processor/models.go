package processor

// Transaction is a struct that holds the summary information for a transaction
type Transaction struct {
	TotalBalance  float64        `json:"total_balance"`
	TotalDebit    float64        `json:"total_debit"`
	TotalCredit   float64        `json:"total_credit"`
	Transactions  map[string]int `json:"transactions"`
	NumberDebits  int            `json:"number_debits"`
	NumberCredits int            `json:"number_credits"`
	DebitAmount   float64        `json:"debit_amount"`
	CreditAmount  float64        `json:"credit_amount"`
	Owner         string         `json:"owner"`
}

// Txn is a struct that holds the information for a transaction
type Txn struct {
	User string  `json:"user"`
	Date string  `json:"date"`
	Txn  float64 `json:"transaction"`
}
