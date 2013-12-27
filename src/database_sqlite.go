package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type User struct {
	Username   string
	Password   string
	Email      string
	CreateTime string
}

func (u User) String() string {
	return fmt.Sprintf("UserName: %s Password: %s CreateTime: %s", u.Username, u.Password, u.CreateTime)
}

/**
 * 检查错误
 */
func checkError(err error) {
	if err != nil {
		// panic(err)
		fmt.Println(err)
	}
}

/**
 * 获取用户列表
 */
func getUserList() {
	db, err := sql.Open("sqlite3", "./user.db")
	checkError(err)
	defer db.Close()
	rows, err := db.Query("select username,departname,created from userinfo order by username desc limit 2000")
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Username, &user.Password, &user.CreateTime)
		checkError(err)
		fmt.Println(user)
	}

}

/**
 * 添加用户
 */
func addUser(u User) error {
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(u.Username, u.Password, u.CreateTime)
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}
func add() {
	u := User{}
	u.Username = fmt.Sprintf("creturn%d", time.Now().Unix())
	u.Password = "jfdklajfkldajsljla"
	u.CreateTime = fmt.Sprintf("%d", time.Now().Unix())
	if err := addUser(u); err != nil {
		fmt.Println(err)
	}
}
func main() {
	// add()
	getUserList()
}
