package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

/**
 * HTTP server
 */
func httpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("server", "API SERVER by:Creturn")
		// writeLog(fmt.Sprintf("%s URL: %s From:%s Time:%s\n", r.Method, r.URL, r.RemoteAddr, time.Now()))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/**
 * 记录日志
 */
func writeLog(msg string) {
	if checkFileIsExist("log.txt") {
		f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		n, err := f.WriteString(fmt.Sprintf("%s\n ", msg))
		if err != nil {
			panic(err)
		}
		fmt.Printf("wirte %dwords", n)
		f.Close()
	} else {
		f, err := os.Create("log.txt")
		if err != nil {
			panic(err)
		}
		n, err := f.WriteString(fmt.Sprintf("%s\n", msg))
		if err != nil {
			panic(err)
		}
		fmt.Printf("wirte %dwords", n)
		f.Close()
	}
}

/**
 * 判断文件是否存在
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

/**
 * Http GET 方式请求
 */
func httpGet(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return fmt.Sprintf("%s", body)
}

/**
 * mysql 测试
 */
func mysql() {
	db, err := sql.Open("mysql", "root:root@/xiaozhan")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	var userNum string
	row := db.QueryRow("select count(*) from gd_member")
	if err := row.Scan(&userNum); err != nil {
		log.Println(err)
	}
	fmt.Printf("----------\nHas %s Numbers member in our site \n----------\n", userNum)
	rows, err := db.Query("select username, password from gd_member order by id desc limit 20")
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var name, pass string
		if err := rows.Scan(&name, &pass); err != nil {
			panic(err)
		}
		fmt.Printf("name: %s pass:%s\n", name, pass)
	}

}
func mapTest() {
	t := make(map[string]int)
	t["a"] = 1
	fmt.Println(t["a"])
	t["a"] = 23
	fmt.Println(t["a"])
	delete(t, "a")
	fmt.Println(t["a"])
	t["a"] = 1000000000000000000
	a, ok := t["a"]
	fmt.Println(a, ok)
}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}
func goTest() {
	var input string
	go say("world")
	fmt.Println("hello")
	fmt.Scanln(&input)
}

/**
 * 入口
 */
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Println("Database test start...\n")
	// goTest()
	mapTest()
	// mysql()
	// httpServer()
	// writeLog("good")

	// fmt.Printf("%s", httpGet("http://www.joy88.com"))
}
