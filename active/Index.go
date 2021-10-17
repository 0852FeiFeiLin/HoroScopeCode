package active

import (
	"HoroScopeCode/entity"
	"HoroScopeCode/page_parse"
	"HoroScopeCode/request"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)
const APIURL = "http://web.juhe.cn/constellation/getAll"
const KEY = "abc8e5d03b9acbe5491877f76470e62d"

func ToIndex(w http.ResponseWriter,r *http.Request)  {
	files, err := template.ParseFiles("./view/index.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	files.Execute(w,nil)

}
//解析用户输入表单
func Index(w http.ResponseWriter,r *http.Request) {
	err := r.ParseForm()
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	//获取前端数据
	consname := r.FormValue("consName")
	type1 := r.FormValue("type")
	/*
		这里应该建立一个判断，，当输入不符合数据类型的时候，程序也能运行
		程序容错性，健壮性，安全性：
				for 循环，容器里面装那些数据。
				那么容器里的数据哪里来，
					1、通过网络请求
					2、数据库中：conf配置表中，通过设置这个去检查数据库，数据库有就返回，没有就返回不存在
	*/
	//if type1 != ()


	fmt.Println(consname)
	fmt.Println(type1)
	//前端数据放入结构体
	userdata := entity.UserData{
		ConsName: consname,
		Type: type1,
	}

	/*http://web.juhe.cn/constellation/getAll?consName=%E7%8B%AE%E5%AD%90%E5%BA%A7&type=today&key=申请的KEY*/


	//这样会报错，因为这里利用了url编码，汉字他需要进行url编码
	//endurl := APIURL + "?consName=" + userdata.ConsName + "&type=" + userdata.Type + "&key=" + KEY

//url参数数据格式化  (url编码格式问题)
	param := url.Values{}
	param.Set("consName",userdata.ConsName)
	param.Set("type",userdata.Type)
	param.Set("key","abc8e5d03b9acbe5491877f76470e62d")
	var Url *url.URL
	Url, err = url.Parse(APIURL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery =  param.Encode()

	//方法2：url.QueryEscape()   汉字编码成url
	//url.QueryUnEscape()  url转为原文
	//url.path()   对路径进行处理

	//发送请求，返回响应数据
	response, err := request.Request("GET", Url.String(), nil)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("请求星座运势api失败")
	}
	fmt.Println(string(response))

	//如果输入的是year年运势的话，就有不同的操作
	if type1== "year" {
		Data2 := page_parse.Parse_Json2(response,)  //还是结构体数据
		fmt.Println("反序列化数据是：",Data2)
		//把结构体数据放进map里面，前端模板遍历
		data := map[string]interface{}{
			"Reason" : Data2,
		}
		files, err := template.ParseFiles("./view/horo.html")
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("跳转页面失败")
			return
		}
		files.Execute(w,data)//把结构体写在页面上
		return
	}

	//解析response返回的数据(反序列化)
	Data := page_parse.Parse_Json(response)  //还是结构体数据

	fmt.Println("反序列化数据是：",Data)

	//把结构体数据放进map里面，前端模板遍历
	data := map[string]interface{}{
		"Reason" : Data,
	}

	files, err := template.ParseFiles("./view/horo.html")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("跳转页面失败")
		return
	}
	files.Execute(w,data)//把结构体写在页面上
}