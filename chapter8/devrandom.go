package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

func main() {
	var seed int64
	binary.Read(rand.Reader, binary.LittleEndian, &seed)
	fmt.Println(seed)
}
