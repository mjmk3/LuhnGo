package payload

type CardValidationResponse struct {
	Valid    bool   `json:"valid"`
	CardType string `json:"cardType"`
}
