package hello

// #include "hello.h"
import "C"

// Hello calls C code hello() func
func Hello() {
	C.hello()
}
