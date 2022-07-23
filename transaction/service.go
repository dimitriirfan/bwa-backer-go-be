package transaction

import (
	"backer/campaign"
	"errors"
)

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionByUserID(input GetUserTransactionsInput) ([]Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

func NewService(repository Repository, campaignRepository campaign.Repository) *service {
	return &service{
		repository:         repository,
		campaignRepository: campaignRepository,
	}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {

	var transactions []Transaction

	campaign, err := s.campaignRepository.FindByCampaignID(input.ID)

	if err != nil {
		return transactions, err
	}

	if input.User.ID != campaign.User.ID {
		return transactions, errors.New("unauthorized")
	}

	transactions, err = s.repository.GetTransactionsByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionByUserID(input GetUserTransactionsInput) ([]Transaction, error) {

	var transactions []Transaction

	transactions, err := s.repository.GetTransactionsByUserID(input.User.ID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
