package main

import (
	"fmt"
)

var text string

func main() {

    testfunc()
	fmt.Println(text)
}

func testfunc() {
	text = "Hello World!"
}
