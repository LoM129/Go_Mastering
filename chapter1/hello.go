package main

import "fmt"
import "net"



func main() {
	fmt.Println("Hello, World!")

/*	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}
		if i == 95 {
			break
		}
	fmt.Print(i, " ")
	}
	*/

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Interface: %v\n", i.Name)
		byName, err := net.InterfaceByName(i.Name)
		if err != nil {
			fmt.Println(err)
		}
		addresses, err := byName.Addrs()
		for k, v := range addresses {
			fmt.Printf("Interface Address #%v: %v\n", k, v.String())
		}
		fmt.Println()
	}	
	
}