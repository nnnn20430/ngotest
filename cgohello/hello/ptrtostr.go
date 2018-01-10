package hello

// #include "hello.h"
import "C"
import "fmt"

func PtrToStr() {
	var ret *C.char = C.hello_string()
	fmt.Printf(C.GoString(ret))
}
