package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmmount      int    `json:"goal_amount"`
	CurrentAmmount   int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmmount:      campaign.GoalAmmount,
		CurrentAmmount:   campaign.CurrentAmmount,
		Slug:             campaign.Slug,
		ImageURL:         "",
	}

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}
	for _, campaign := range campaigns {
		campaignsFormatter = append(campaignsFormatter, FormatCampaign(campaign))
	}

	return campaignsFormatter
}
