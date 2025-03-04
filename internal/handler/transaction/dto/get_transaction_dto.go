package dto

type QueryTransaction struct {
	Page       int    `form:"page"`
	Limit      int    `form:"limit"`
	Search     string `form:"search"`
	ConsumerId string `form:"consumer_id"`
	OrderBy    string `form:"order_by"`
	Direction  string `form:"direction"`
}
