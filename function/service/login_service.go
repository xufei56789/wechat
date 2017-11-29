package service

import (
	_"time"
	"fmt"
	"time"
	http "xufei/wechat/function"
	"xufei/wechat"
)


func GetUUID() (string, error){
	uuidPara := map[string]string{
		"appid":     "wx782c26e4c19acffb",
		"fun":       "new",
		"lang":      "zh_CN",
		"timeStamp": ""}
	uuidPara["timeStamp"]	= fmt.Sprintf("%d", time.Now().Unix())
	var result string = "?"
	for key, value := range uuidPara {
		result += fmt.Sprintf("%s=%s&",key, value)
	}
	resp , err:= http.Get(wechat.UUID_URL + result)
	if err != nil {
		panic("访问微信服务器失败")
	}
	defer resp.Body.Close()


	fmt.Println(resp)







	return "2", nil
}


