package main

func init() {
	fulfillCardTypeBankNumWithBankServiceType()
}

func fulfillCardTypeBankNumWithBankServiceType() {
	var abank ABank
	var bbank BBank

	// TODO:
	PaymentOptionsMap[PaymentOptionKeyCreator("4", "242 10")] = abank.Free
	PaymentOptionsMap[PaymentOptionKeyCreator("4", "242 20")] = abank.Gold
	PaymentOptionsMap[PaymentOptionKeyCreator("4", "242 30")] = abank.Platinum
	PaymentOptionsMap[PaymentOptionKeyCreator("5", "242 10")] = abank.Free
	PaymentOptionsMap[PaymentOptionKeyCreator("5", "242 20")] = abank.Gold
	PaymentOptionsMap[PaymentOptionKeyCreator("5", "242 30")] = abank.Platinum
	PaymentOptionsMap[PaymentOptionKeyCreator("5", "567 10")] = bbank.Free
	PaymentOptionsMap[PaymentOptionKeyCreator("5", "567 20")] = bbank.Platinum
}

func SeeThePaymentOptions(cardType string, bankNum string, price float64) string {
	key := PaymentOptionKeyCreator(cardType, bankNum)

	handler, ok := PaymentOptionsMap[key]
	if !ok {
		return "Unsupported Card"
	}

	return handler(price)
}
