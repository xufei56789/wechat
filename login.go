package wechat

import (
	"xufei/wechat/function"
	"fmt"
	"net/url"
)

//获取web微信uuid
func (w *Wechat) GetUUid() (string , error) {
	//获取时间搓19位纳秒
	sTime := function.GetNewTime()
	//窃取0-13位
	sTime = function.RegexpString(sTime,0,13)
	//获取uuid
	resp , err := w.httpx.Get("https://login.wx.qq.com/jslogin?appid=wx782c26e4c19acffb&redirect_uri=https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage&fun=new&lang=zh_CN&_="+sTime)
	if err != nil && resp.Status != 200 {
		return "" , err
	}
	//[window.QRLogin.code = 200; window.QRLogin.uuid = "QbsAGoam5Q==";]
	//获取code 22-25位
	code := string(resp.Body[22:25])
	if code != "200" {
		return "" , ErrCodeNo200
	}
	//获取uuid 50-62位
	uuid := string(resp.Body[50:62])
	return uuid , nil
}

//获取二维码
func (w *Wechat) GetQrcide(uuid string) ([]byte ,error) {
	resp , err := w.httpx.Get("https://login.weixin.qq.com/qrcode/"+uuid)
	if err != nil && resp.Status != 200 {
		return []byte("") , err
	}
	return resp.Body , nil
}
//个人信息
func (w *Wechat) GetLoginPara() map[string]string{
	var loginPara = map[string]string{
		"loginicon":"true",
		"tip":       "0",
		"uuid":		"",
		"R":		"",
		"TimeStamp": "",
	}
	return loginPara
}

//验证是否登录
func (w *Wechat) CheckLogin(uuid string){
	sTime := function.GetNewTime()
	sTime = function.RegexpString(sTime,0,13)
	paraMap := w.GetLoginPara()
	paraMap["uuid"] = uuid
	paraMap["TimeStamp"] = sTime
	paraMap["R"] = reverseString(sTime)
	uPar := w.FormatParams(paraMap)
	resp , err :=w.httpx.Get("https://login.weixin.qq.com/cgi-bin/mmwebwx-bin/login"+uPar)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

//格式数据
func (w *Wechat) FormatParams(values ...interface{}) string{
	var result = "?"
	maap := values[0].(map[string]string)
	for key, value := range maap {
		if key != "" && value != "" {
			result += fmt.Sprintf("%s=%s&", key, url.QueryEscape(value))
		}
	}
	return result
}

func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from + 1, to - 1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}



