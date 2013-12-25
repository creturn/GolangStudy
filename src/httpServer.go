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
func main() {
	runtime.SetCPUProfileRate(runtime.NumCPU())
	startServer()
}
