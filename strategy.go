package main

import "fmt"

func PurchaseCreditCardStatus(price float64, ratio float64, numberOfInstallment int) string {
	interestedPrice := price + price*ratio
	installmentPricePerMonth := interestedPrice / float64(numberOfInstallment)

	return fmt.Sprintf("Taksit Sayısı: %d, Toplam Geri Ödeme: %.02f, Taksit Tutarı: %.02f",
		numberOfInstallment, interestedPrice, installmentPricePerMonth)
}
