package main

import (
	"xufei/wechat/function/service"
	"fmt"
)

func main()  {
	uuId,_ := service.GetUUID()
	fmt.Println(uuId)
}
