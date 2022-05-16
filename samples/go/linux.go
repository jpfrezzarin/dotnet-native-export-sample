package main

import "fmt"

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <stdio.h>

int MyLibSumBridge(void *f, int a, int b) { // wrapper function
    //description: ((return_data_type (*)(input_data_type_1, ...))bridge_input_function_pointer) (bridge_input_value_1, ...)
    return ((int (*)(int, int))f)(a, b);
}
*/
import "C"

func main() {

	fmt.Println("Hello, world!")

	handle := C.dlopen(C.CString("./output/MyNativeLibraryNE.so"), C.RTLD_LAZY)
	sum_ptr := C.dlsym(handle, C.CString("MyLibSum"))
	result := C.MyLibSumBridge(sum_ptr, 1, 2)

	fmt.Printf("The sum result is %p\n", result)
}
