package domain

type GetTransfersResponse struct {
	Transfers []Transfer
}

type GetTransferByID struct {
	Transfer *Transfer
}

type Transfer struct {
	ID       int    `json:"id"`
	IDFrom   int    `json:"id_from"`
	IDTo     int    `json:"id_to"`
	Currency string `json:"currency"`
	Value    int    `json:"value"`
}

type TransferFilter struct {
	IDFrom int `form:"id_from"`
	IDTo   int `form:"id_to"`
}
