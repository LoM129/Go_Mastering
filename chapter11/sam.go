package main

import (
	"fmt"
	"os"
	"strconv"
)

func sam(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + sam(n-1)
	}
}

func main() {
	argument := os.Args
	if len(argument) != 2 {
		fmt.Println("This code needs 1 argument!")
		return
	}
	n, err := strconv.Atoi(argument[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	samm := sam(n)

	fmt.Println("sam:", samm)

}
