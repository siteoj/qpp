package main

import (
	"fmt"
	"syscall"
	"time"
)

func main() {
	test := syscall.NewLazyDLL("DllTestDef.dll")
	add := test.NewProc("add")
	outRand := make([]byte, 2)
	a := 1
	b := 4
	sum, _, _ := add.Call(uintptr(a), uintptr(b))
	fmt.Printf("sum:%d\noutRand:%s\n", sum, outRand)
	for{
		time.Sleep(time.Millisecond*100)
	}
}