package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func PostToZibal(path string, requestData interface{}) (PaymentResponse, error) {
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("failed to marshal request data: %w", err)
	}

	url := "https://gateway.zibal.ir/" + path
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	var response PaymentResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return response, nil
}

func RequestResult(result string) string {
	switch result {
	case "100":
		return "با موفقیت تایید شد."
	case "102":
		return "merchant یافت نشد."
	case "103":
		return "Mamerchant غیرفعالrch"
	case "104":
		return "merchant نامعتبر"
	case "201":
		return "قبلا تایید شده."
	case "105":
		return "amount بایستی بزرگتر از 1,000 ریال باشد."
	case "106":
		return "callbackUrl نامعتبر می‌باشد. (شروع با http و یا https)"
	case "113":
		return "amount مبلغ تراکنش از سقف میزان تراکنش بیشتر است."

	}
	return "خطا در پرداخت"
}

func VerifyResult(result string) string {
	switch result {
	case "100":
		return "با موفقیت تایید شد."
	case "102":
		return "merchant یافت نشد."
	case "103":
		return "merchant غیر فعال"
	case "104":
		return "merchant نامعتبر"
	case "201":
		return "قبلا تایید شده."
	case "202":
		return "سفارش پرداخت نشده یا ناموفق بوده است."
	case "203":
		return "trackId نامعتبر می‌باشد."
	}
	return "خطا در تایید پرداخت"
}
