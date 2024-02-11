package main

import (
	"fmt"
	"strconv"
)

func main() {
	handlePaymentRequest()
	handlePaymentVerification("your-track-id")
}

func handlePaymentRequest() {
	merchant := "zibal"

	requestData := PaymentRequest{
		Merchant:    merchant,
		CallbackURL: "https://your-domain/callbackUrl",
		Description: "golang package",
		Amount:      10000,
	}

	response, err := PostToZibal("v1/request", requestData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(RequestResult(strconv.Itoa(response.Result)))
}

func handlePaymentVerification(trackID string) {
	merchant := "zibal"

	requestData := VerificationRequest{
		Merchant: merchant,
		TrackID:  trackID,
	}

	response, err := PostToZibal("v1/verify", requestData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(VerifyResult(strconv.Itoa(response.Result)))
}
