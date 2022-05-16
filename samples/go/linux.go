package main

import "fmt"

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <stdio.h>

// syntax: typedef return_type (*func_pointer_type_name)(input_type_1, input_type_2, ...)
typedef int (*MyLibSumType)(int, int); // function pointer type
typedef int (*MyLibSubType)(int, int); // function pointer type

int MyLibSumBridge(void *f, int a, int b) { // wrapper function
    return ((MyLibSumType)f)(a, b);
}
int MyLibSubBridge(void *f, int a, int b) { // wrapper function
    return ((MyLibSubType)f)(a, b);
}
*/
import "C"

func main() {

	fmt.Println("Hello, world!")

	handle := C.dlopen(C.CString("./output/MyNativeLibraryNE.so"), C.RTLD_LAZY)

	sum_ptr := C.dlsym(handle, C.CString("MyLibSum"))
	sub_ptr := C.dlsym(handle, C.CString("MyLibSub"))

	sum_result := C.MyLibSumBridge(sum_ptr, 1, 2)
	sub_result := C.MyLibSubBridge(sub_ptr, 4, 2)

	fmt.Printf("The result of sum is: %d\n", sum_result)
	fmt.Printf("The result of sub is: %d\n", sub_result)
}
