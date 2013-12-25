package main

import (
	"fmt"
	"runtime"
)

type Message struct {
	id      int
	content string
	msgType int
}

func (msg Message) String() string {
	return fmt.Sprintf("ID:%d content:%s Type:%d", msg.id, msg.content, msg.msgType)
}

func main() {
	msg := Message{1, "test a content", 0}
	fmt.Printf("The msg:\n%s\n", msg)
	fmt.Printf("Now have %d task running!", runtime.NumGoroutine())
}
