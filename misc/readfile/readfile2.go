package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var (err error; f *os.File; fi os.FileInfo; buf []byte)
	if len(os.Args) < 2 {
		fmt.Printf("No file specified as first parameter.\n")
		os.Exit(1)
	}
	f, err = os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fi, _ = f.Stat()
	buf = make([]byte, fi.Size())
	io.ReadFull(f, buf)
	f.Close()
	fmt.Printf("%v", string(buf))
	buf = nil
}
