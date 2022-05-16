package main

import (
	"fmt"
	"syscall"
)

func main() {

	fmt.Println("Hello, world!")

	var mod = syscall.NewLazyDLL("output/MyNativeLibraryNE.dll")

	var sum_proc = mod.NewProc("MyLibSum")
	var sub_proc = mod.NewProc("MyLibSub")

	sum_result, _, _ := sum_proc.Call(uintptr(1), uintptr(2))
	sub_result, _, _ := sub_proc.Call(uintptr(4), uintptr(2))

	fmt.Printf("The result of sum is: %d\n", sum_result)
	fmt.Printf("The result of sub is: %d\n", sub_result)
}
