package main

import (
	"fmt"
	//"github.com/gin-gonic/gin"
)

// a := 1
var a = 1

func main() {
	//fmt.Println(a)

	b := []byte{}

	b = make([]byte, 1)
	fmt.Println(b)
	b[0] = 1
	b = append(b, 1)
	fmt.Println(b)
	fmt.Println("hello gin")

	hash := make(map[string]int, 1)
	hash["a"] = 1
	hash["b"] = 2
	fmt.Println(hash, " ", len(hash))

	t1 := []int{1, 2, 3}

	fmt.Println(t1)
}
