package notifier

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ftResp struct {
	Errno   int    `json:"errno"`
	Errmsg  string `json:"errmsg"`
	Dataset string `json:"dataset"`
}

func Ft(key string, title string, username string) {
	fmt.Println(title)
	if key == "" {
		return
	}
	resp, err := http.Get("https://sc.ftqq.com/" + key + ".send?text=" + url.QueryEscape(username + title))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	r := ftResp{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		// handle error
	}
	if r.Errno == 0 && r.Errmsg == "success" {
		fmt.Println("微信通知发送成功")
	} else {
		fmt.Println("微信通知发送失败：" + r.Errmsg)
	}
}
