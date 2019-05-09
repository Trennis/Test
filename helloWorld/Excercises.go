package main

import "fmt"

type myType int

var x myType

func main() {
	fmt.Println(x)

	fmt.Printf("%T\n", x)

	x = 43
	fmt.Println(x)
	fmt.Printf("%#v", x)

}
