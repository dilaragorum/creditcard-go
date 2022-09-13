package main

import "fmt"

type BBank struct{}

func (b BBank) Free(price float64) string {
	if price < 1000 {
		firstOption := PurchaseCreditCardStatus(price, 5.6/100, 3)

		secondOption := PurchaseCreditCardStatus(price, 8.9/100, 6)

		thirdOption := PurchaseCreditCardStatus(price, 0.11, 9)

		return fmt.Sprintf("%s \n %s \n %s", firstOption, secondOption, thirdOption)
	}

	return PurchaseCreditCardStatus(price, 0, 1)
}

func (b BBank) Platinum(price float64) string {
	if price < 1500 {
		firstOption := PurchaseCreditCardStatus(price, 0, 3)

		secondOption := PurchaseCreditCardStatus(price, 5.9, 6)

		thirdOption := PurchaseCreditCardStatus(price, 8, 9)

		result := fmt.Sprintf("%s \n %s \n %s", firstOption, secondOption, thirdOption)

		return result
	}

	return PurchaseCreditCardStatus(price, 0, 1)
}
