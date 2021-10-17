package entity

type Data struct {
	Error_code int `json:"error_code"`
	Reason Reason `json:"reason"`
}
