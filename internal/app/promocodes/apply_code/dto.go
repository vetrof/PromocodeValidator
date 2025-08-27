package apply_code

type Input struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Output struct {
	Code       string `json:"code"`
	Exists     bool   `json:"exists"`
	OnTime     bool   `json:"onTime"`
	Applied    bool   `json:"applied"`
	AppliedNow bool   `json:"appliedNow"`
}
