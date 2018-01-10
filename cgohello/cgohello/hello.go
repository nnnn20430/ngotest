package cgohello

// #include "hello.h"
import "C"

func World() {
	C.hello()
}
