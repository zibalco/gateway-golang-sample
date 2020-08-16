package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
	"net/http"
	// "encoding/json"
)

func postToZibal(path string, parameters string) string{
	var jsonStr = []byte(parameters)
	var url = "https://gateway.zibal.ir/" + path

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}


func requestResult(result string) string{
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

func verifyResult(result string) string{
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
		case "202": 
			return "سفارش پرداخت نشده یا ناموفق بوده است."
		case "203": 
			return "trackId نامعتبر می‌باشد."
	}
	return "خطا در تایید پرداخت"
}