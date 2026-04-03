package main

import (
	"blogx_server/global"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetAccessToken(code string) (accessToken string, err error) {
	qq := global.Config.QQ

	baseUrl, _ := url.Parse("https://graph.qq.com/oauth2.0/token")
	p := url.Values{}
	p.Add("grant_type", "authorization_code")
	p.Add("client_id", qq.AppID)
	p.Add("client_secret", qq.AppKey)
	p.Add("code", code)
	p.Add("redirect_uri", qq.Redirect)
	p.Add("need_openid", "1")
	baseUrl.RawQuery = p.Encode()
	response, err := http.Get(baseUrl.String())
	if err != nil {
		return
	}
	byteData, _ := io.ReadAll(response.Body)
	fmt.Println(string(byteData))
	return
}

func main() {
}
