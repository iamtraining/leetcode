package main

import (
	"fmt"
	"time"
)

var (
	fooC = make(chan bool)
	barC = make(chan bool)
)

func main() {
	n := 5
	go foo(n)
	go bar(n)
	time.Sleep(2 * time.Second)
}

func foo(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("foo")
		fooC <- true
		<-barC
	}
}

func bar(n int) {
	for i := 0; i < n; i++ {
		<-fooC
		fmt.Print("bar")
		barC <- true
	}
}
