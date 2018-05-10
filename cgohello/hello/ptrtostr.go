package hello

// #include "hello.h"
import "C"
import "fmt"

// PtrToStr converts a C string to a Go string using C.GoString()
func PtrToStr() {
	var ret *C.char = C.hello_string()
	fmt.Printf(C.GoString(ret))
}
