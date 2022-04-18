package main

import (
	"fmt"
	"time"
)

func main() {
	out1 := make(chan string)
	go process("order", out1)
	for msg := range out1 {
		fmt.Println(msg)
	}
}

func process(item string, out chan string) {
	defer close(out) // close the channel when stop using it
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second / 2)
		out <- item
	}
}
