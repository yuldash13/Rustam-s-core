package db

type Status string

const (
	Success  Status = "success"
	Canceled Status = "cancel"
)

type Transfer struct {
	ID             int
	IDFrom         int
	IDTo           int
	Currency       string
	Value          int
	OperationState Status
}

type TransferFilter struct {
	IDFrom int
	IDTo   int
}
