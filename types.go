package main

type PaymentRequest struct {
	Merchant    string `json:"merchant"`
	CallbackURL string `json:"callbackUrl"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type PaymentResponse struct {
	Result  int    `json:"result"`
	TrackID string `json:"trackId"`
}
