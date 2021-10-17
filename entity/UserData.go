package entity

//用户输入的信息
type UserData struct {
	ConsName string `form:"consName"`
	Type string `form:"type"`
}
