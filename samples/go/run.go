package main

import (
	"fmt"
	"syscall"
)

func main() {

	fmt.Println("Hello, world!")

	var mod = syscall.NewLazyDLL("output/MyNativeLibraryNE.dll")
	var proc = mod.NewProc("Sum")

	ret, _, _ := proc.Call(
		uintptr(1),
		uintptr(2))

	fmt.Printf("Return: %d\n", ret)
}
