package bpackage

import (
	"fmt"
	"mymodule/apackage"
)

func init() {
	fmt.Println("init() b")
}
func FromB() {
	fmt.Println("fromB()")
	apackage.FromA()
}
