package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

//获取天气api地址
const APIURL = "http://api.openweathermap.org/data/2.5/weather?q=Shanghai&mode=xml"

type Current struct {
	XMLName     xml.Name          `xml:"current"`
	City        CityEntity        `xml:"city"`
	Temperature TemperatureEntity `xml:"temperature"`
	LastUpdate  LastUpdateEntity  `xml:"lastupdate"`
}
type TemperatureEntity struct {
	Value string `xml:"value,attr"`
	Min   string `xml:"min,attr"`
	Max   string `xml:"max,attr"`
	Unit  string `xml:"unit,attr"`
}
type CityEntity struct {
	Id      string      `xml:"id,attr"`
	Name    string      `xml:"name,attr"`
	Coor    CoordEntity `xml:"coord"`
	Country string      `xml:"country"`
	Sun     SunEntity   `xml:"sun"`
}
type CoordEntity struct {
	Lon string `xml:"lon,attr"`
	Lat string `xml:"lat,attr"`
}
type SunEntity struct {
	Rise string `xml:"rise,attr"`
	Set  string `xml:"set,attr"`
}
type LastUpdateEntity struct {
	Value string `xml:"value,attr"`
}

/**
 * 工程师
 */
type Engineer struct {
	XMLName   xml.Name `xml:"engineer"`
	Comment   string   `xml:",comment"` //",comment" 指的是此字段为注释
	Id        string   `xml:"id"`
	Name      string   `xml:"name"`
	Email     string   `xml:"email"`
	Count     int      `xml:"-"`            //"-" 标识不解析次字段
	SkillData []Skill  `xml:"skills>skill"` //这里需要注意的是skills>skill 的意思是父节点叫做skills包含子节点skill
}

/**
 * 技能
 */
type Skill struct {
	Skillname string `xml:"name"`
	Skillyear string `xml:"year"`
	Project   string `xml:"project,omitempty"` //"omitempty" 表示字段为空时候不解析
}

/**
 *	输出xml
 */
func xmlEncode() {
	u := Engineer{Id: "1", Name: "creturn", Email: "master@creturn.com", Comment: "工程师资料"}
	u.SkillData = append(u.SkillData, Skill{Skillname: "php", Skillyear: "4", Project: "More"})
	u.SkillData = append(u.SkillData, Skill{Skillname: "javascript", Skillyear: "4"})
	u.SkillData = append(u.SkillData, Skill{Skillname: "python", Skillyear: "1"})

	data, err := xml.MarshalIndent(&u, "", "	") //第一个参数是其实间距，第二个是指子元素和父元素的距离
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(xml.Header), string(data))
}

/**
 * 解析xml数据
 */
func xmlDecode() {
	resp, err := http.Get(APIURL)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
	weather := Current{}
	xml.Unmarshal(content, &weather)
	fmt.Println(weather)
}
func main() {
	// xmlEncode()
	xmlDecode()
}
