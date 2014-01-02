package main

import (
	"bufio"
	"fmt"
	// "io/ioutil"
	"net"
	"os"
)

/**
 * IRCServer
 */
type IrcServer struct {
	Domain  string
	Port    int
	Channel string
}

/**
 * IRC客户端
 */
type IrcClient struct {
	Nick     string
	UserName string
	RealName string
	Con      net.Conn
	Server   IrcServer
	Pwrite   chan string
}

/**
 * 启动
 */
func (irc *IrcClient) Run() {
	fmt.Println("Client start....")
	fmt.Println("Now connect to ", irc.Server.Domain, "on the Channel:", irc.Server.Channel)
	irc.connect()
	go irc.loopReceive()
	go irc.loopWirte()
	irc.Register()
	irc.Join(irc.Server.Channel)

}

/**
 * 链接
 */
func (irc *IrcClient) connect() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", irc.Server.Domain, irc.Server.Port))
	if err != nil {
		irc.log("Resolve err:" + err.Error())
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		irc.log("connect err:" + err.Error())
	}
	irc.Con = conn
	irc.Pwrite = make(chan string, 1024)
}

/**
 * 发送消息
 */
func (irc *IrcClient) loopWirte() {
	for {
		msg, ok := <-irc.Pwrite
		if !ok || msg == "" || irc.Con == nil {
			break
		}
		_, err := irc.Con.Write([]byte(msg))
		if err != nil {
			irc.log("socket write err:" + err.Error())
			break
		}
	}
}

/**
 * 断开连接
 */
func (irc *IrcClient) Disconnect() {
	irc.Con.Close()
	fmt.Println("disconnect")
}

/**
 * 循环接收消息
 */
func (irc *IrcClient) loopReceive() {
	bf := bufio.NewReaderSize(irc.Con, 512)
	for {
		msg, err := bf.ReadString('\n')
		if err != nil {
			irc.log("ReceiveError:" + err.Error())
		}
		fmt.Println("Receive Msg:")
		fmt.Println(msg)
	}
}

/**
 * 注册登陆
 */
func (irc *IrcClient) Register() {
	irc.SendRawF("NICK %s\r\n", irc.Nick)
	irc.SendRawF("USER %s 0.0.0.0 0.0.0.0 :%s\r\n", irc.UserName, irc.RealName)
}

/**
 * 加入channel
 */
func (irc *IrcClient) Join(channel string) {
	irc.SendRawF("JOIN %s\r\n", channel)
}

/**
 * 发送消息，通过通道（多线程）
 */
func (irc *IrcClient) SendRawF(m string, data ...interface{}) {
	irc.Pwrite <- fmt.Sprintf(m, data) + "\r\n"
}

/**
 * 记录错误日志
 */
func (irc *IrcClient) log(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	ircs := IrcServer{Domain: "chat.freenode.net", Port: 6667, Channel: "#creturnTest"}

	client := &IrcClient{Server: ircs, UserName: "CreturnBot", Nick: "BotReturn"}
	client.Run()

	for {
		var msg string
		fmt.Scanln(&msg)
		fmt.Println("you type:" + msg)
	}
	defer client.Disconnect()
	os.Exit(1)
}
