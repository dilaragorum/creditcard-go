package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strings"
)

type BankHandler func(price float64) string

var (
	PaymentOptionsMap       = make(map[string]BankHandler)
	PaymentOptionKeyCreator = func(cardType, bankNum string) string {
		return fmt.Sprintf("%s.%s", cardType, bankNum)
	}
)

func main() {
	e := echo.New()

	e.POST("/options", CalculateThePaymentOptions)
	e.POST("/payment", PaymentType)

	e.Logger.Fatal(e.Start(":8080"))
}

func CalculateThePaymentOptions(c echo.Context) error {
	var request PaymentOptionRequest

	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cardType := request.GetCardType()
	bankNum := request.GetBankNum()
	price := request.Price

	log.Infof("cardType %s, bankNum %s", cardType, bankNum)

	options := SeeThePaymentOptions(cardType, bankNum, price)

	return c.String(http.StatusOK, options)
}

func PaymentType(c echo.Context) error {
	var request PaymentOptionRequest

	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	cardType := request.GetCardType()
	bankNum := request.GetBankNum()
	key := PaymentOptionKeyCreator(cardType, bankNum)

	if _, ok := PaymentOptionsMap[key]; ok {
		if cardType == "4" {
			marshal, _ := json.Marshal(request)
			buffer := bytes.NewBuffer(marshal)
			response, err := http.Post("https://trendyol-case.mocklab.io/v1/pay", "application/json", buffer)
			if err != nil {
				return err
			}
			defer response.Body.Close()

			return c.JSON(http.StatusCreated, response.Body)
		}
		bankNum = strings.ReplaceAll(bankNum, " ", "")
		url := fmt.Sprintf("https://typay.com/v1/pay/%s/%s", bankNum, uuid.NewString())
		return c.Redirect(http.StatusFound, url)
	}

	return c.String(http.StatusNotFound, "Unsupported Card")
}
