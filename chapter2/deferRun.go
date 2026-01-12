package main

import (
	"fmt"
)

func deferRun() (res int) {
  var num = 1
  	defer func() {
		res += 2*1 + 3
	}()
  num = 2
  return num
}

func main() {

	n := deferRun()
	fmt.Println("This is test defer: " , n)
}