package main

import (
	"fmt"
)

type ABank struct{}

func (a ABank) Free(price float64) string {
	if price <= 100 {
		return PurchaseCreditCardStatus(price, 0.05, 3)
	}

	return PurchaseCreditCardStatus(price, 0, 1)
}

func (a ABank) Gold(price float64) string {
	if price <= 100 {
		return PurchaseCreditCardStatus(price, 0, 3)
	} else if price > 100 && price <= 500 {
		return PurchaseCreditCardStatus(price, 0.05, 6)
	}

	return PurchaseCreditCardStatus(price, 0, 1)
}

func (a ABank) Platinum(price float64) string {
	if price <= 500 {
		firstOption := PurchaseCreditCardStatus(price, 0, 3)
		secondOption := PurchaseCreditCardStatus(price, 0, 6)

		result := fmt.Sprintf("%s \n %s", firstOption, secondOption)

		return result
	}

	return PurchaseCreditCardStatus(price, 0.05, 9)
}
