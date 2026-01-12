package main

import (
	"fmt"
	"mymodule/apackage"
	"mymodule/bpackage"
)

func init() {
	fmt.Println("init() manyInit")
}

func main() {
	fmt.Println("Using aPackage!")
	apackage.FromA()
	bpackage.FromB()
	fmt.Println(apackage.MyConstant)
}
