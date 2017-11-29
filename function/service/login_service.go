package service

import (
	_"time"
	"fmt"
	"time"
	http "xufei/wechat/function"
	"xufei/wechat"
)

const (
	APPID 	= "wx782c26e4c19acffb"
	FUN		= "new"
	Lang 	= "zh_CN"
	TimeStamp = ""
)

func GetUUID() (string, error){
	uuidPara := map[string]string{
		APPID:     "wx782c26e4c19acffb",
		FUN:       "new",
		Lang:      Lang,
		TimeStamp: ""}
	uuidPara[TimeStamp]	= fmt.Sprintf("%d", time.Now().Unix())
	fmt.Println(uuidPara)



	http.Get(wechat.UUID_URL + "s")







	return "2", nil
}


