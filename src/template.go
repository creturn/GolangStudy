package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type User struct {
	Username   string
	Password   string
	Email      string
	CreateTime string
}
type UserList struct {
	MemberList []User
}

/**
 * 登陆
 */
func login(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tpl/login.html")
	t.Execute(w, nil)
}
func userList(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tpl/user_list.html")
	List := getUserList()
	// fmt.Println(list)
	t.Execute(w, List)
}

/**
 * 获取用户列表
 */
func getUserList() UserList {
	var users []User
	db, err := sql.Open("mysql", "root:root@/xiaozhan")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	rows, err := db.Query("select username,password from gd_member order by id desc limit 20")
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Username, &user.Password)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
		// checkError(err)
		//		fmt.Println(user)
	}
	return UserList{users}

}

/**
 *首页
 */
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

/**
 * 启动服务
 */
func startServer() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/user", userList)
	http.ListenAndServe(":8080", nil)

}
func main() {
	startServer()
}
