package main

import (
	"fmt"
	"os"
)

func main() {
	var name string = "Joe Jonas"
	var age int = 27

	fmt.Println(name, "is", age)

	changeName(name)

	fmt.Println(name, "is", age)

	realChangeName(&name)

	fmt.Println(name, "is", age)
}

func changeName(name string) {
	name = "Doof Jordan"

	fmt.Println(name, "has changed names")
}

func realChangeName(name *string) string {
	*name = os.Getenv("USER")

	fmt.Println(*name, "has changed names")

	return *name
}
