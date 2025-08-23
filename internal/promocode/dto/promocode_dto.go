package dto

type ValidateInput struct {
	Code string `json:"code"`
}

type ValidateOutput struct {
	Valid   bool   `json:"valid"`
	Error   string `json:"error,omitempty"`
	Success bool   `json:"success"`
}
