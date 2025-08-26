package dto

type ValidateInput struct {
	Code string `json:"code"`
}

type ValidateOutput struct {
	Code  string `json:"code"`
	Valid bool   `json:"valid"`
}
