package main

import (
	"fmt"
	// "github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

/**
 * 获取地址内容
 */
func getUrlContent(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	//设置头部标识为iphone
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 6_1_3 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10B329 Safari/8536.25")
	client := http.Client{}
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}

/**
 * 获取贴吧内容列表
 */
func getTiebaList(tName string) {
	tbUrl := fmt.Sprintf("http://tieba.baidu.com/f?kw=%s&lp=1501&mo_device=1", tName)
	content := getUrlContent(tbUrl)
	fmt.Println(content)
}
func main() {
	// getUrlContent("http://tieba.baidu.com/f?kw=大主宰&lp=1501&mo_device=1")
	getTiebaList("大主宰")
}
