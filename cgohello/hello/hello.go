package hello

// #include "hello.h"
import "C"

func Hello() {
	C.hello()
}
