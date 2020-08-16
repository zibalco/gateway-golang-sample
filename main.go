package main

import (
    "fmt"
	"encoding/json"
	"bytes"
)

func main() {

	// Request Sample Code
	var merchant = "zibal"

	data := `{
        "merchant" : "` + merchant + `",
        "callbackUrl" : "https://your-domain/callbackUrl",
        "description" : "golang package",
        "amount" : 10000
	}`

	var result = postToZibal("v1/request", data)
	// Map result to a struct to easily access parameters
	var structResult map[string]interface{}
	br := bytes.NewReader([]byte(result))
	decodedJson := json.NewDecoder(br)
	decodedJson.UseNumber()
	err := decodedJson.Decode(&structResult)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// Access response parameters
	var resultNumber = structResult["result"]
	var trackId = structResult["trackId"]
	// Change response parameters types to string
	trackIdStringValue := fmt.Sprint(trackId)
	resultStringValue := fmt.Sprint(resultNumber)

	// Print readable messages based on response result code
	fmt.Println(requestResult(resultStringValue))


	// Verify Sample Code
	data = `{
        "merchant" : "` + merchant + `",
        "trackId" : ` + trackIdStringValue + `
	}`

	result = postToZibal("v1/verify", data)

	// Map result to a struct to easily access parameters
	var structVerify map[string]interface{}
	br = bytes.NewReader([]byte(result))
	decodedJson = json.NewDecoder(br)
	decodedJson.UseNumber()
	err = decodedJson.Decode(&structVerify)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	// Access response parameters
	var verifyResultCode = structVerify["result"]

	// Change response parameters types to string
	verifyResultCodeStringValue := fmt.Sprint(verifyResultCode)

	// Print readable messages based on response result code
	fmt.Println(verifyResult(verifyResultCodeStringValue))
	
}