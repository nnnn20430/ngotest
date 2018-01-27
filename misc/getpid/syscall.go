package main

import "syscall"
import "fmt"

func main() {
	var pid uintptr
	pid, _, _ = syscall.Syscall(syscall.SYS_GETPID, 0, 0, 0)
	fmt.Printf("%v\n", int(pid))
}
