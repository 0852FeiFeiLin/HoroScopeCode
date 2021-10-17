package request

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)
//模拟客户端发送请求返回响应
func Request(method,url string,body io.Reader)(re []byte,err error){
	//创建一个客户端
	client := http.Client{
		//设置请求超时时间
		Timeout: 30 * time.Second,
	}
	//创建一个请求对象
	request, err := http.NewRequest(method, url, body)
	if err != nil{
		fmt.Println(err.Error())
		return nil,err
	}
	//客户端发送请求
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if response.StatusCode != 200{
		fmt.Println(err.Error())
		return
	}
	//读取响应体数据
	bytes, err := ioutil.ReadAll(response.Body)


	return bytes,nil
}
