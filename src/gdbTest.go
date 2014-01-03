package main

import (
	"fmt"
	"time"
)

func counting(c chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		c <- i
	}
	close(c)
}

/**
 * gdb 调试程序的时候生成代码需要加入参数
 * go build -gcflags "-N -l" gdbTest.go
 *
 * run 运行程序
 * b [number] 在多少行下断点
 * info locals 查看当前变量信息
 */
func main() {
	msg := "Start Main"
	fmt.Println(msg)
	bus := make(chan int)
	go counting(bus)
	for count := range bus {
		fmt.Println("count:", count)
	}
}
