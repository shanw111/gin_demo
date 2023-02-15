package main

import (
	"fmt"
	"unsafe"
	//"github.com/gin-gonic/gin"
)

// a := 1
var a = 1

type print interface {
	print_v()
	//print_k()
}
type pen struct {
	name string
}

func (r *pen) print_v() {
	fmt.Println(r.name)
}

//func (r *pen)print_k(){
//	fmt.Println(r.name)
//}

func main() {

	var v print = &pen{" pen"}
	v.print_v()

	i := 0
	var c = 'a'
	p := &c
	fmt.Println(unsafe.Sizeof(i), " , ", unsafe.Sizeof(p))

	fmt.Println(unsafe.Sizeof("UESTC"))
	fmt.Println(unsafe.Sizeof("电子科技大学"))
	fmt.Println(len("电子科技大学"))
	fmt.Println(len("UESTC"))

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
