package cgohello

// #include "hello.h"
import "C"
import "fmt"
import "unsafe"

func PtrToArr() {
	var ret *C.char = C.hello_string()
	fmt.Printf(string((*(*[]byte)(unsafe.Pointer(&ret)))[:14]))
}
