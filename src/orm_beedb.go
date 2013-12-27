package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

/**
 * 字段首字母必须大写
 * go 语法规定，首字母大写才是Public
 * 不然在其他包中就无法使用例如beedb的orm
 */
type News struct {
	Id      int //`beedb:"PK"`
	Title   string
	Content string
	Time    string
}

func (n News) String() string {
	return fmt.Sprintf("ID: %d , Title: %s, Content: %s, Time: %s", n.Id, n.Title, n.Content, n.Time)
}

/**
 * 获取一条新闻
 */
func getOneNews(id int) {
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	orm := beedb.New(db)
	var OneData News
	orm.Where("id=?", id).Find(&OneData)
	fmt.Println(OneData)
}

/**
 * 获取所有新闻
 */
func getAllNews() {
	start := time.Now()
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	orm := beedb.New(db)
	var AllData []News
	orm.FindAll(&AllData)
	for i := 0; i < len(AllData); i++ {
		fmt.Println(AllData[i])
	}
	fmt.Println("耗时：", time.Since(start))
}

/**
 * 获取所有新闻非orm
 * 经测试rom还是比较耗时的
 * 不过有点就是开发效率
 */
func getAllByBase() {
	start := time.Now()
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from news")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var one News
		rows.Scan(&one.Id, &one.Title, &one.Content, &one.Time)
		fmt.Println(one)
	}
	fmt.Println("耗时：", time.Since(start))
}
func main() {
	// getAllNews()
	getAllByBase()
}
