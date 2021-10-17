package entity

type Year struct {
	Name string `json:"name"`
	Date string `json:"date"`
	Year int `json:"year"`
	Mima Mima `json:"mima"`
	Career[] string `json:"career"`
	Love[] string `json:"love"`
	Health[] string `json:"health"`
	Finance[] string `json:"finance"`
	LuckeyStone[] string `json:"luckey_stone"`
	Future string `json:"future"`
	Resultcode string `json:"resultcode"`
	Error_code int `json:"error_code"`
}