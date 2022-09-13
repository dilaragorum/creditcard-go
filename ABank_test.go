package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestABank_Free(t *testing.T) {
	t.Run("Success ABank_Free - When price equal and lower than 100", func(t *testing.T) {
		type testCase struct {
			price          float64
			expectedResult string
		}

		ABank := ABank{}

		testCases := []testCase{
			{price: 100, expectedResult: "Taksit Sayısı: 3, Toplam Geri Ödeme: 105.00, Taksit Tutarı: 35.00"},
			{price: 90, expectedResult: "Taksit Sayısı: 3, Toplam Geri Ödeme: 94.50, Taksit Tutarı: 31.50"},
		}

		for _, test := range testCases {
			expectedResult := test.expectedResult
			receivedResult := ABank.Free(test.price)

			assert.Equal(t, expectedResult, receivedResult)
		}
	})

	t.Run("Success ABank_Free - When price higher than 100", func(t *testing.T) {
		type testCase struct {
			price          float64
			expectedResult string
		}

		ABank := ABank{}

		testCases := []testCase{
			{price: 120, expectedResult: "Taksit Sayısı: 1, Toplam Geri Ödeme: 120.00, Taksit Tutarı: 120.00"},
			{price: 125, expectedResult: "Taksit Sayısı: 1, Toplam Geri Ödeme: 125.00, Taksit Tutarı: 125.00"},
		}

		for _, test := range testCases {
			expectedResult := test.expectedResult
			receivedResult := ABank.Free(test.price)

			assert.Equal(t, expectedResult, receivedResult)
		}
	})
}

func TestABank_Gold(t *testing.T) {
	t.Run("Success ABank_Gold - When price equal and lower than 100", func(t *testing.T) {
		type testCase struct {
			price          float64
			expectedResult string
		}

		ABank := ABank{}

		testCases := []testCase{
			{price: 100, expectedResult: "Taksit Sayısı: 3, Toplam Geri Ödeme: 100.00, Taksit Tutarı: 33.33"},
			{price: 90, expectedResult: "Taksit Sayısı: 3, Toplam Geri Ödeme: 90.00, Taksit Tutarı: 30.00"},
		}

		for _, test := range testCases {
			expectedResult := test.expectedResult
			receivedResult := ABank.Gold(test.price)

			assert.Equal(t, expectedResult, receivedResult)
		}
	})

	t.Run("Success ABank_Gold - When price higher than 100 and lower than 500", func(t *testing.T) {
		type testCase struct {
			price          float64
			expectedResult string
		}

		ABank := ABank{}

		testCases := []testCase{
			{price: 102, expectedResult: "Taksit Sayısı: 6, Toplam Geri Ödeme: 107.10, Taksit Tutarı: 17.85"},
			{price: 200, expectedResult: "Taksit Sayısı: 6, Toplam Geri Ödeme: 210.00, Taksit Tutarı: 35.00"},
		}

		for _, test := range testCases {
			expectedResult := test.expectedResult
			receivedResult := ABank.Gold(test.price)

			assert.Equal(t, expectedResult, receivedResult)
		}
	})

	t.Run("Success ABank_Gold - When price higher than 500", func(t *testing.T) {
		type testCase struct {
			price          float64
			expectedResult string
		}

		ABank := ABank{}

		testCases := []testCase{
			{price: 501, expectedResult: "Taksit Sayısı: 1, Toplam Geri Ödeme: 501.00, Taksit Tutarı: 501.00"},
			{price: 1000, expectedResult: "Taksit Sayısı: 1, Toplam Geri Ödeme: 1000.00, Taksit Tutarı: 1000.00"},
		}

		for _, test := range testCases {
			expectedResult := test.expectedResult
			receivedResult := ABank.Gold(test.price)

			assert.Equal(t, expectedResult, receivedResult)
		}
	})
}

func TestABank_Platinum(t *testing.T) {
	t.Run("Success ABank_Platinum - When price lower than 500", func(t *testing.T) {
		type testCase struct {
			price          float64
			expectedResult string
		}

		ABank := ABank{}

		testCases := []testCase{
			{price: 500, expectedResult: "Taksit Sayısı: 3, Toplam Geri Ödeme: 500.00, Taksit Tutarı: 166.67 \n Taksit Sayısı: 6, Toplam Geri Ödeme: 500.00, Taksit Tutarı: 83.33"},
			{price: 100, expectedResult: "Taksit Sayısı: 3, Toplam Geri Ödeme: 100.00, Taksit Tutarı: 33.33 \n Taksit Sayısı: 6, Toplam Geri Ödeme: 100.00, Taksit Tutarı: 16.67"},
		}

		for _, test := range testCases {
			expectedResult := test.expectedResult
			receivedResult := ABank.Platinum(test.price)

			assert.Equal(t, expectedResult, receivedResult)
		}
	})

	t.Run("Success ABank_Platinum - When price higher than 500", func(t *testing.T) {
		type testCase struct {
			price          float64
			expectedResult string
		}

		ABank := ABank{}

		testCases := []testCase{
			{price: 502, expectedResult: "Taksit Sayısı: 9, Toplam Geri Ödeme: 527.10, Taksit Tutarı: 58.57"},
			{price: 1000, expectedResult: "Taksit Sayısı: 9, Toplam Geri Ödeme: 1050.00, Taksit Tutarı: 116.67"},
		}

		for _, test := range testCases {
			expectedResult := test.expectedResult
			receivedResult := ABank.Platinum(test.price)

			assert.Equal(t, expectedResult, receivedResult)
		}
	})

}
