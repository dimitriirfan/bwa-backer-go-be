package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}

	return formatter

}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	formatter := []CampaignTransactionFormatter{}

	for _, transaction := range transactions {
		formatter = append(formatter, FormatCampaignTransaction(transaction))
	}

	return formatter
}
