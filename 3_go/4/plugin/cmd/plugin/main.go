package main

import "C"
import (
	"fmt"
)

//export PrintLine
func PrintLine(cstr *C.char) {
	fmt.Println(C.GoString(cstr))
}

func main() {
	//	just empty main file to build pluginn
}
