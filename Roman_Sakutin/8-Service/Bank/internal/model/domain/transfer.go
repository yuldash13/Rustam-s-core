package domain

type Status string

const (
	Success Status = "success"
	Canceled Status = "cancel"
)

type Transfer struct {
	ID             int    `json:"id"`
	IDFrom         int    `json:"id_from"`
	IDTo           int    `json:"id_to"`
	TransDate      string `json:"trans_date"`
	Currency       string `json:"currency"`
	Value          int    `json:"value"`
	OperationState Status
}

type TransferFilter struct {
	IDFrom int `form:"id_from"`
	IDTo   int `form:"id_to"`
}
