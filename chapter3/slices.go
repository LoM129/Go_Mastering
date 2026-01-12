package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
person string
height int
weight int
}

func printSlice(x []int) {
	for _, number := range x {
		fmt.Print(number, "  ")
	}
	fmt.Println()
}

func main(){

// first part : cap expend
	integer := make([]int, 4)
	integer2 := new([]int)
	fmt.Println(integer2)
	for i := 0; i < len(integer); i++ {
		fmt.Println(integer[i])
	}


	aSlice := []int{-1, 0, 4}
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -100)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))

//second part : memory maping

	fmt.Println()
	fmt.Println("This is another example of slice.")
	s1 := make([]int, 5)
	reSlice := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reSlice)
	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(s1)
	fmt.Println(reSlice)

//third part : sorting

	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Mihalis", 180, 90})
	mySlice = append(mySlice, aStructure{"Bill", 134, 45})
	mySlice = append(mySlice, aStructure{"Marietta", 155, 45})
	mySlice = append(mySlice, aStructure{"Epifanios", 144, 50})
	mySlice = append(mySlice, aStructure{"Athina", 134, 40})
	fmt.Println("0:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)
	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)
}