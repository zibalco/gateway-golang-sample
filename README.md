# Zibal Payment Gateway Golang Code Sample

You can use the function "PostToZibal" from "functions.go" to create a payment request and verify the payments.
Functions "RequestResult" and "VerifyResult" can be called for printing readable messages based on the result code.

# Request payment sample
```go
// Request Sample Code
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

```


# Verify payment sample
```go
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

```

