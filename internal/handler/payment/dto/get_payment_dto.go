package dto

type QueryPayment struct {
	Page           int    `form:"page"`
	Limit          int    `form:"limit"`
	Search         string `form:"search"`
	ConsumerId     string `form:"consumer_id"`
	ContractNumber string `form:"contract_number"`
	Status         string `form:"status"`
	OrderBy        string `form:"order_by"`
	Direction      string `form:"direction"`
}
