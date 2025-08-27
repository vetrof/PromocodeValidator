package apply_code

type PromocodeInput struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type PromocodeOutput struct {
	Code       string `json:"code"`
	Exists     bool   `json:"exists"`
	OnTime     bool   `json:"onTime"`
	Applied    bool   `json:"applied"`
	AppliedNow bool   `json:"appliedNow"`
}
