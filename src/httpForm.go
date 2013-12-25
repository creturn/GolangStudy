package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

/**
 * 默认登陆页面
 */
func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./tpl/login.html")
	t.Execute(w, nil)
}

/**
 * 登陆
 */
func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //用于解析表单数据
	for key, val := range r.Form {
		fmt.Fprintf(w, "user:%s pass:%s\n", key, val[0])
		fmt.Fprintf(w, "user:%s pass:%s\n", key, r.Form.Get(key))

	}
	fmt.Fprintln(w, "token is :"+r.FormValue("token"))
	fmt.Fprint(w, "logining")
}

/**
 * 原生post数据测试
 * 根据HTTP协议原生没有经过解析的数据是在body里面
 * 所以再解析一些post数据的时候，例如微信API里面post
 * 过来的数据用From是无法获取的
 */
func post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(w, string(body))
}

/**
 * 上传文件
 * 服务端接收客户端上传的文件
 */
func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		t, _ := template.ParseFiles("./tpl/upload.html")
		t.Execute(w, nil)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, hanle, err := r.FormFile("upfile")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", hanle.Header)
		f, err := os.OpenFile("./upload/"+hanle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

/**
 * 客户端模拟用户上传文件
 */
func clientUploadFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWrite := multipart.NewWriter(bodyBuf)

	fileWrite, err := bodyWrite.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error write to buffer")
		return err
	}
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error openfile")
		return err
	}
	size, err := io.Copy(fileWrite, fh)
	if err != nil {
		return err
	}
	contentType := bodyWrite.FormDataContentType()
	bodyWrite.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(respBody))
	fmt.Println("size:" + size)
	return nil
}

/**
 * 启动web服务
 */
func startServer() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/post", post)
	http.HandleFunc("/upload", uploadFile)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	startServer()
}
