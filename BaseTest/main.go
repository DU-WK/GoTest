package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Println("hello world!")

	inta := 10
	stra := "ceshi"
	strb := "测试"

	fmt.Println(unsafe.Sizeof(inta))
	fmt.Println(reflect.TypeOf(stra[0]))
	fmt.Println(reflect.TypeOf(strb))

	fmt.Println(stra[0], string(stra[0]))

	fmt.Printf("%d, %c\n", stra[0], stra[0])

	fmt.Println("len(stra)", len(stra))
	fmt.Println("len(strb)", len(strb))
}
