package transaction

import (
	"backer/campaign"
	"time"
)

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionFormatter struct {
	ID        int                            `json:"id"`
	Amount    int                            `json:"amount"`
	Status    string                         `json:"status"`
	CreatedAt time.Time                      `json:"created_at"`
	Campaign  campaign.CampaignUserFormatter `json:"campaign"`
}

type TransactionFormatter struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
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

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
	}

	campaignFormatter := campaign.CampaignUserFormatter{
		Name:     transaction.Campaign.Name,
		ImageURL: "",
	}

	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	formatter := []UserTransactionFormatter{}

	for _, transaction := range transactions {
		formatter = append(formatter, FormatUserTransaction(transaction))
	}

	return formatter
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{
		ID:         transaction.ID,
		CampaignID: transaction.CampaignID,
		UserID:     transaction.UserID,
		Amount:     transaction.Amount,
		Status:     transaction.Status,
		Code:       transaction.Status,
		PaymentURL: transaction.PaymentURL,
	}

	return formatter
}
