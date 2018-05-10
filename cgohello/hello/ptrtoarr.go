package hello

// #include "hello.h"
import "C"
import "fmt"
import "unsafe"

// PtrToArr converts a C string to a byte slice and then to a Go string
// uses fixed string length of 14
func PtrToArr() {
	var ret *C.char = C.hello_string()
	fmt.Printf(string((*(*[]byte)(unsafe.Pointer(&ret)))[:14]))
}
