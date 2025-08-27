package valid_code

type Output struct {
	Code    string `json:"code"`
	Exists  bool   `json:"exists"`
	OnTime  bool   `json:"onTime"`
	Applied bool   `json:"applied"`
}
