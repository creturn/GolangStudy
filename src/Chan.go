package main

import (
	"fmt"
	"os"
	"time"
)

func say(name string, c chan string) {
	for {
		time.Sleep(5 * 1000 * time.Millisecond)
		if s := <-c; s != "exit" {
			fmt.Printf("%s listion you say: %s time:%s\n", name, s, time.Now())
		} else {
			fmt.Println(name, "exit")
			return
		}
	}

}
func main() {
	var input string
	fmt.Println("start chan")
	ch := make(chan string, 10)
	go say("fly", ch)
	go say("xiu", ch)
	for {
		if fmt.Scan(&input); input != "quit" {
			fmt.Println("The god say:", input)
			ch <- input
		} else {
			fmt.Println("Good bye.")
			return
		}
	}
}
