package main

import (
    "fmt"
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

	fmt.Println(name, "has changed")
}

func realChangeName(name *string) string {
	*name = "Doof Jordan"

	fmt.Println(*name, "has changed")

	return *name
}
