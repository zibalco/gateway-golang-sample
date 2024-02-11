package main

import (
	"log"
	"strconv"
)

func main() {
	log.Println(log.LstdFlags | log.Lshortfile)
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
		log.Println("Error:", err)
		return
	}

	log.Println(RequestResult(strconv.Itoa(response.Result)))
}

func handlePaymentVerification(trackID string) {
	merchant := "zibal"

	requestData := VerificationRequest{
		Merchant: merchant,
		TrackID:  trackID,
	}

	response, err := PostToZibal("v1/verify", requestData)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	log.Println(VerifyResult(strconv.Itoa(response.Result)))
}
