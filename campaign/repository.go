package campaign

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByCampaignID(campaignID int) (Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByCampaignID(campaignID int) (Campaign, error) {
	var campaign Campaign

	err := r.db.Where("ID = ?", campaignID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
