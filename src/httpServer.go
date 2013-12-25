package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

/**
 * Server默认请求地址
 */
func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Helllo wellcome!")
}

/**
 * 获取作者信息
 */
func author(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "The author: creturn blog: www.creturn.com")
}

/**
 * 获取状态信息
 */
func stat(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "It works ok!")
}

/**
 * URL 路由绑定
 */
func route() {
	http.HandleFunc("/", start)
	http.HandleFunc("/author", author)
	http.HandleFunc("/status", stat)
}

/**
 * 启动web Server
 */
func startServer() {
	route()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("started server!")
}

type MyMux struct {
}

/**
 * 路由分发器
 * 注意此接口名称必须是ServeHTTP
 * 因为默认路由调用的就是ServeHTTP这个接口
 */
func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		start(w, r)
	case "/author":
		author(w, r)
	case "/stat":
		stat(w, r)
	default:
		http.NotFound(w, r)
	}
}

/**
 * 自定义路由分发测试
 */
func customRoute() {
	//注意这里的MyMux 必须实现ServerHTTP接口才可以使用
	mux := &MyMux{}
	http.ListenAndServe(":8089", mux)
}

/**
 * 入口函数
 */
func main() {
	runtime.SetCPUProfileRate(runtime.NumCPU())
	//一般http server 处理
	go startServer()
	//自定义路由分发处理
	customRoute()

}
