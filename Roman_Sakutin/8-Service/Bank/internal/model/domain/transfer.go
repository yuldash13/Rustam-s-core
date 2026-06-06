package domain

type Status string

const (
	Success  Status = "success"
	Canceled Status = "cancel"
)

type Transfer struct {
	ID             int64
	IDFrom         int64
	IDTo           int64
	Currency       string
	Value          int64
	OperationState Status
}

type TransferFilter struct {
	IDFrom int64
	IDTo   int64
}
