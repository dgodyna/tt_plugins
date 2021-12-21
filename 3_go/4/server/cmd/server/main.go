package main

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}
#cgo LDFLAGS: -L${SRCDIR}
#cgo LDFLAGS: -lplugin
#include <plugin.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

func PrintLine(s string) {
	cstr := C.CString(s)
	defer C.free(unsafe.Pointer(cstr))
	C.PrintLine(cstr)
}

func main() {
	PrintLine("test")
}
