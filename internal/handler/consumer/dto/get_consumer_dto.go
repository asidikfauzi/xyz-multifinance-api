package dto

type QueryConsumer struct {
	Page      int    `form:"page"`
	Limit     int    `form:"limit"`
	Search    string `form:"search"`
	OrderBy   string `form:"order_by"`
	Direction string `form:"direction"`
	Paginate  string `form:"paginate"`
}
