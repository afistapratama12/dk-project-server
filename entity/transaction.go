package entity

type Transaction struct {
	Id           int `json:"id"`
	FromId       int `json:"from_id"`
	ToId         int `json:"to_id"`
	SASBalance   int `json:"sas_balance"`
	ROBalance    int `json:"ro_balance"`
	MoneyBalance int `json:"money_balance"`
}

type TransInput struct {
	FromId       int `json:"from_id" binding:"required"`
	ToId         int `json:"to_id" binding:"required"`
	SASBalance   int `json:"sas_balance"`
	ROBalance    int `json:"ro_balance"`
	MoneyBalance int `json:"money_balance"`
}

type NewDownlineInput struct {
	UplineId     int `json:"upline_id"`
	MoneyBalance int `json:"money_balance"`
}
