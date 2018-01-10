package cgohello

// #include "hello.h"
import "C"
import "fmt"
import "unsafe"

func PtrArithmetic() {
	var ret *C.char = C.hello_string()
	var start uintptr = uintptr(unsafe.Pointer(ret))
	var size uintptr = unsafe.Sizeof(C.char(0))
	var char byte
	for i := 0; true; i++ {
		char = (*(*byte)(unsafe.Pointer(start + size*uintptr(i))))
		if char == 0 {
			break
		}
		fmt.Printf(string(char))
	}
}
