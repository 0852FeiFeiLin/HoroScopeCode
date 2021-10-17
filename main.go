package main

import (
	"HoroScopeCode/active"
	"fmt"
	"net/http"
)

//星座运势api代码实现，还差分开实现，today，tomorrow...
/*
	思路：
		1、首先前端用户输入，输入自己的星座，和运势类型
		2、获取用户输入的内容，发起请求
		3、获取返回的response，
		4、返回的response反序列化，
		5、响应数据存入结构体
		6、显示在前端
*/
func main() {
	//配置路径
	http.HandleFunc("/",active.ToIndex)
	http.HandleFunc("/horoscope",active.Index)
	//获取响应数据


	//开启监听，启动服务器
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println("服务器创建失败")
		fmt.Println(err.Error())
		return
	}

	fmt.Println("服务器启动成功")

}
