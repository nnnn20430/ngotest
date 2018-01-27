// +build ignore

package main

import "fmt"
import "os"
import "syscall"

func main() {
	var pid int
	if len(os.Args) > 1 && os.Args[1] == "forked" {
		fmt.Printf("forked!\n")
		os.Exit(0)
	}
	pid, _ = syscall.ForkExec(os.Args[0], []string{os.Args[0], "forked"}, &syscall.ProcAttr{Files: []uintptr{0, 1, 2}})
	syscall.Wait4(pid, nil, 0, nil)
}
