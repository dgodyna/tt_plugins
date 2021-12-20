package main

import "C"
import (
	"fmt"
)

//export PrintLine
func PrintLine(s string) {
	fmt.Println(s)

}

func main() {
	//	just empty main file to build pluginn
}
