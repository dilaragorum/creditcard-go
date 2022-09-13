package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPurchaseCreditCardStatus(t *testing.T) {
	type testCase struct {
		price               float64
		ratio               float64
		numberOfInstallment int
		expectedResult      string
	}
	testCases := []testCase{

		{price: 100, ratio: 0.05, numberOfInstallment: 3, expectedResult: "Taksit Sayısı: 3, Toplam Geri Ödeme: 105.00, Taksit Tutarı: 35.00"},
		{price: 125, ratio: 0.06, numberOfInstallment: 6, expectedResult: "Taksit Sayısı: 6, Toplam Geri Ödeme: 132.50, Taksit Tutarı: 22.08"},
		{price: 1250, ratio: 0.05, numberOfInstallment: 5, expectedResult: "Taksit Sayısı: 5, Toplam Geri Ödeme: 1312.50, Taksit Tutarı: 262.50"},
	}

	for _, test := range testCases {
		expectedResult := test.expectedResult
		receivedResult := PurchaseCreditCardStatus(test.price, test.ratio, test.numberOfInstallment)
		assert.Equal(t, expectedResult, receivedResult)

		fmt.Printf("case: price = %f ratio = %f expectedResult = %s  recievedResult = %s", test.price, test.ratio, expectedResult, receivedResult)
	}
}
