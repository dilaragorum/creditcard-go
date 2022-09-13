package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBBank_Free(t *testing.T) {
	t.Run("Success BBank_Free - When price lower than 1000", func(t *testing.T) {
		type testCase struct {
			price          float64
			expectedResult string
		}
		testCases := []testCase{
			{100, "Taksit Sayısı: 3, Toplam Geri Ödeme: 105.60, Taksit Tutarı: 35.20 \n Taksit Sayısı: 6, Toplam Geri Ödeme: 108.90, Taksit Tutarı: 18.15 \n Taksit Sayısı: 9, Toplam Geri Ödeme: 111.00, Taksit Tutarı: 12.33"},
			{550, "Taksit Sayısı: 3, Toplam Geri Ödeme: 580.80, Taksit Tutarı: 193.60 \n Taksit Sayısı: 6, Toplam Geri Ödeme: 598.95, Taksit Tutarı: 99.83 \n Taksit Sayısı: 9, Toplam Geri Ödeme: 610.50, Taksit Tutarı: 67.83"},
		}

		BBank := BBank{}

		for _, test := range testCases {
			expectedResult := test.expectedResult
			receivedResult := BBank.Free(test.price)

			assert.Equal(t, expectedResult, receivedResult)
		}
	})

	t.Run("Success BBank_Free - When price higher than 1000", func(t *testing.T) {
		type testCase struct {
			price          float64
			expectedResult string
		}
		testCases := []testCase{
			{1020, "Taksit Sayısı: 1, Toplam Geri Ödeme: 1020.00, Taksit Tutarı: 1020.00"},
			{2000, "Taksit Sayısı: 1, Toplam Geri Ödeme: 2000.00, Taksit Tutarı: 2000.00"},
		}

		BBank := BBank{}

		for _, test := range testCases {
			expectedResult := test.expectedResult
			receivedResult := BBank.Free(test.price)

			assert.Equal(t, expectedResult, receivedResult)
		}
	})

}

func TestBBank_Platinum(t *testing.T) {
	t.Run("Success BBank_Platinum - When price lower than 1500", func(t *testing.T) {
		type testCase struct {
			price          float64
			expectedResult string
		}

		bBank := BBank{}

		testCases := []testCase{
			{price: 400, expectedResult: "Taksit Sayısı: 3, Toplam Geri Ödeme: 400.00, Taksit Tutarı: 133.33 \n Taksit Sayısı: 6, Toplam Geri Ödeme: 2760.00, Taksit Tutarı: 460.00 \n Taksit Sayısı: 9, Toplam Geri Ödeme: 3600.00, Taksit Tutarı: 400.00"},
			{price: 500, expectedResult: "Taksit Sayısı: 3, Toplam Geri Ödeme: 500.00, Taksit Tutarı: 166.67 \n Taksit Sayısı: 6, Toplam Geri Ödeme: 3450.00, Taksit Tutarı: 575.00 \n Taksit Sayısı: 9, Toplam Geri Ödeme: 4500.00, Taksit Tutarı: 500.00"},
		}

		for _, test := range testCases {
			expectedResult := test.expectedResult
			actualResult := bBank.Platinum(test.price)

			assert.Equal(t, expectedResult, actualResult)
		}
	})

	t.Run("Success BBank_Platinum - When price higher than 1500", func(t *testing.T) {
		type testCase struct {
			price          float64
			expectedResult string
		}

		bBank := BBank{}

		testCases := []testCase{
			{price: 1500, expectedResult: "Taksit Sayısı: 1, Toplam Geri Ödeme: 1500.00, Taksit Tutarı: 1500.00"},
			{price: 2000, expectedResult: "Taksit Sayısı: 1, Toplam Geri Ödeme: 2000.00, Taksit Tutarı: 2000.00"},
		}

		for _, test := range testCases {
			expectedResult := test.expectedResult
			actualResult := bBank.Platinum(test.price)

			assert.Equal(t, expectedResult, actualResult)
		}
	})
}
