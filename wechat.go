package wechat

import "github.com/changx123/httpx"

type Wechat struct {
	//httpx指针
	httpx *httpx.Httpx
}

//获取新的httpx指针对象
func getNewHttpx() *httpx.Httpx {
	hx := httpx.NewHttpx()
	//使用cookie Jar容器自动保存cookie信息
	hx.SetAutoSaveCookie(true)
	//阻止302跳转
	hx.SetRedirect(0)
	return hx
}

//设置代理ip
func (w *Wechat) SetProxy(url string) {
	w.httpx.SetProxy(url)
}
