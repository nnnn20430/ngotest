package main

import (
	"reflect"
	"runtime/debug"
)

func ptrFree(V interface{}) {
	Vpv := reflect.ValueOf(V)
	Vpt := Vpv.Type()
	if Vpt.Kind() == reflect.Ptr {
		Vt := Vpt.Elem()
		Vv := reflect.Indirect(Vpv)
		Zv := reflect.Zero(Vt)
		Vv.Set(Zv)
		debug.FreeOSMemory()
	} else {
		panic("Free(): parameter is not a pointer")
	}
}
