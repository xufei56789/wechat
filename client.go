package wechat

import (
	"github.com/changx123/httpx"
	"fmt"
)

type Client struct {
	appid string
	secret string
	token string
}

func NewClient(appid, secret, token string) *Client {
	return &Client{
		appid: appid,
		secret: secret,
		token: token,
	}
}

func (c *Client) Access (url string) error {
	agent,err := httpx.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(agent)
	return  nil
}
