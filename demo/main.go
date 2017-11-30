package main

import (
	"xufei/wechat"
	"fmt"
	"os"
)
func main(){
	var wechat = &wechat.Wechat{}
	wechat.Init()
	uuid, err := wechat.GetUUid()
	if err != nil {
		fmt.Println(err)
		return
	}
	img, err := wechat.GetQrcide(uuid)
	if err != nil{
		fmt.Println(err)
		return
	}
	file, err := os.OpenFile("./二维码.png", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	file.Write(img)

	//轮询判断是否已经登录过了
	for {
		wechat.CheckLogin(uuid)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

