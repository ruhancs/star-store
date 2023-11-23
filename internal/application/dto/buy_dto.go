package dto

type InputBuyUseCaseDto struct {
	ClientID   string  `json:"client_id"`
	Value      float32 `json:"value"`
	CardNumber string  `json:"card_number"`
}

type OutputBuyUseCaseDto struct {
	ClientID         string  `json:"client_id"`
	ClientName       string  `json:"client_name"`
	TotalToPay       float32 `json:"total_to_pay"`
	CreditCardNumber string  `json:"credit_card_number"`
}
