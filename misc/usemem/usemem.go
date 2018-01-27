package main

import (
	"reflect"
	"runtime/debug"
	"syscall"
	"time"
	"fmt"
)

func free(V interface{}) {
	Vpv := reflect.ValueOf(V)
	Vpt := Vpv.Type()
	if Vpt.Kind() == reflect.Ptr {
		Vt := Vpt.Elem()
		Vv := reflect.Indirect(Vpv)
		Zv := reflect.Zero(Vt)
		Vv.Set(Zv)
		debug.FreeOSMemory()
	} else {
		panic("free(): parameter is not a pointer")
	}
}

func main() {
	var (buf []byte; sysinfo syscall.Sysinfo_t)
	syscall.Sysinfo(&sysinfo)
	fmt.Printf("allocating %v bytes\n", sysinfo.Freeram/2)
	buf = make([]byte, sysinfo.Freeram/2)
	for i := range buf {
		buf[i] = 'a'
	}
	time.Sleep(5 * time.Second)
	free(&buf)
	fmt.Printf("freed\n")
	time.Sleep(5 * time.Second)
}
