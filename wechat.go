package wechat

import (
	"github.com/changx123/httpx"
	"net/http"
)

type Wechat struct {
	//httpx指针
	httpx *httpx.Httpx
}

//用户浏览器类型(防封设置)
var UserAgent = []string{"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0"}

//获取新的httpx指针对象
func getNewHttpx() *httpx.Httpx {
	hx := httpx.NewHttpx()
	//使用cookie Jar容器自动保存cookie信息
	hx.SetAutoSaveCookie(true)
	//阻止302跳转
	hx.SetRedirect(0)
	//设置协议头
	var h *http.Header
	h.Set("User-Agent", UserAgent[0])
	h.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	h.Set("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	h.Set("Connection", "keep-alive")
	h.Set("Upgrade-Insecure-Requests", "1")
	hx.SetHeader(h)
	return hx
}

//设置代理ip
func (w *Wechat) SetProxy(url string) {
	w.httpx.SetProxy(url)
}
