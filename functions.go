package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	resultCode, err := strconv.Atoi(result)
	if err != nil {
		return "خطا در پرداخت"
	}

	message, found := ResultMessages[resultCode]
	if !found {
		return "خطا در پرداخت"
	}

	return message
}

func VerifyResult(result string) string {
	resultCode, err := strconv.Atoi(result)
	if err != nil {
		return "خطا در تایید پرداخت"
	}

	message, found := ResultMessages[resultCode]
	if !found {
		return "خطا در تایید پرداخت"
	}

	return message
}
