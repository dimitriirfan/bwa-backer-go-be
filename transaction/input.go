package transaction

import "backer/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type GetUserTransactionsInput struct {
	User user.User
}
