package main

type PaymentOptionRequest struct {
	CardNum string  `json:"cardNumber"`
	Price   float64 `json:"price"`
}

func (p *PaymentOptionRequest) GetCardType() string {
	return p.CardNum[0:1]
}

func (p *PaymentOptionRequest) GetBankNum() string {
	return p.CardNum[1:7]
}
