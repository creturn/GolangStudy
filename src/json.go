package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// "net/http"
)

/**
 * pm2.5获取地址
 * @type {String}
 */
const APIURL = "http://www.pm25.in/api/querys/aqi_ranking.json?token=kvqaTz7tzLtyXJffzhDu"

/**
 * 定义终端颜色输出
 * @type {String}
 */
const (
	NONE         = "\033[m"
	RED          = "\033[0;32;31m"
	LIGHT_RED    = "\033[1;31m"
	GREEN        = "\033[0;32;32m"
	LIGHT_GREEN  = "\033[1;32m"
	BLUE         = "\033[0;32;34m"
	LIGHT_BLUE   = "\033[1;34m"
	DARY_GRAY    = "\033[1;30m"
	CYAN         = "\033[0;36m"
	LIGHT_CYAN   = "\033[1;36m"
	PURPLE       = "\033[0;35m"
	LIGHT_PURPLE = "\033[1;35m"
	BROWN        = "\033[0;33m"
	YELLOW       = "\033[1;33m"
	LIGHT_GRAY   = "\033[0;37m"
	WHITE        = "\033[1;37m"
	CLOSE_ATTR   = "\033[m"
)

/**
 * 空气质量
 */
type AirQuality struct {
	Aqi       int    `json:"aqi"`
	Area      string `json:"area"`
	Co        int    `json:"co"`
	Co_24h    int    `json:"co_24h"`
	Pm2_5     int    `json:"pm2_5`
	Pm2_5_24h int    `json:"pm2_5_24h"`
	Quality   string `json:"quality"`
	Level     string `json:"level"`
	Time      string `json:"time_point,omitempty"`
}
type AirList []AirQuality
type Member struct {
	QQNum     int    `json:"uin"`
	Nick      string `json:"nick"`
	Iscreator int
	Ismanager int
}

func (air AirQuality) String() string {
	return fmt.Sprintf("城市: %s Pm2.5: %d", air.Area, air.Pm2_5)
}
func (mem Member) String() string {
	var Nick string
	if mem.Iscreator == 1 {
		Nick = RED + mem.Nick + CLOSE_ATTR
	}
	return fmt.Sprintf("QQ: %d 昵称: %s", mem.QQNum, Nick)
}

/**
 * 输出json
 */
func encodeJson() {
	var airList []AirQuality
	airList = append(airList, AirQuality{Aqi: 222, Area: "陕西", Pm2_5: 10})
	airList = append(airList, AirQuality{Aqi: 400, Area: "北京", Pm2_5: 500})
	airList = append(airList, AirQuality{Aqi: 300, Area: "上海", Pm2_5: 300})
	jsonOut, err := json.Marshal(airList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonOut))

}

/**
 * 解析json
 */
func decodeJson() {
	content, err := ioutil.ReadFile("pm25.json")
	if err != nil {
		fmt.Println(err)
	}
	var list AirList
	json.Unmarshal(content, &list)
	// fmt.Println(list)
	fmt.Println(BLUE, "天气质量排行", CLOSE_ATTR)
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
}
func testColor() {
	fmt.Println(RED, "RED", CLOSE_ATTR)
	fmt.Println(LIGHT_RED, "LIGHT_RED", CLOSE_ATTR)
	fmt.Println(GREEN, "GREEN", CLOSE_ATTR)
	fmt.Println(LIGHT_GREEN, "LIGHT_GREEN", CLOSE_ATTR)
	fmt.Println(BLUE, "BLUE", CLOSE_ATTR)
	fmt.Println(LIGHT_BLUE, "LIGHT_BLUE", CLOSE_ATTR)
	fmt.Println(DARY_GRAY, "DARY_GRAY", CLOSE_ATTR)
	fmt.Println(CYAN, "CYAN", CLOSE_ATTR)
	fmt.Println(LIGHT_CYAN, "LIGHT_CYAN", CLOSE_ATTR)
	fmt.Println(PURPLE, "PURPLE", CLOSE_ATTR)
	fmt.Println(LIGHT_PURPLE, "LIGHT_PURPLE", CLOSE_ATTR)
	fmt.Println(BROWN, "BROWN", CLOSE_ATTR)
	fmt.Println(YELLOW, "YELLOW", CLOSE_ATTR)
	fmt.Println(LIGHT_GRAY, "LIGHT_GRAY", CLOSE_ATTR)
	fmt.Println(WHITE, "GREEN", CLOSE_ATTR)
	fmt.Println(NONE, "NONE", CLOSE_ATTR)
}
func decodeGolangQunMmeber() {
	content, err := ioutil.ReadFile("qun.json")
	if err != nil {
		fmt.Println(err)
	}
	var memberList []Member
	json.Unmarshal(content, &memberList)
	for i := 0; i < len(memberList); i++ {
		fmt.Println(memberList[i])
	}
}
func main() {
	// encodeJson()
	// decodeJson()
	decodeGolangQunMmeber()
}
