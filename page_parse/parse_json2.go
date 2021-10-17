package page_parse

import (
	"HoroScopeCode/entity"
	"encoding/json"
	"fmt"
)

func Parse_Json2(foro []byte) (re entity.Year){

	var data entity.Year
	fmt.Println(foro)
	jsonerr := json.Unmarshal(foro, &data)
	if jsonerr != nil{
		fmt.Println(jsonerr.Error())
		fmt.Println("解码异常")
		return
	}
	if data.Error_code != 0{  //返回码为0 == 成功
		fmt.Println("请求失败",data.Error_code)
	}
	fmt.Println(data)
	return data  //返回解析好的数据
}
