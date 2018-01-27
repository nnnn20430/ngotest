package main

import "fmt"
import "os"

func main() {
	var p *os.Process
	if len(os.Args) > 1 && os.Args[1] == "forked" {
		fmt.Printf("forked!\n")
		os.Exit(0)
	}
	p, _ = os.StartProcess(os.Args[0], []string{os.Args[0], "forked"}, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
	p.Wait()
}
