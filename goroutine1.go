package main

import (
	"fmt"
	"time"
	r "runtime"
)

var f = fmt.Println

func function() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	f()
}

func main() {
//	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
	go function()
	f("Number of Goroutines:", r.NumGoroutine())
	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()
	time.Sleep(1 * time.Second)
	f("Number of Goroutines:", r.NumGoroutine())
}