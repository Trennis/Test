package main

import "fmt"

var xyz int
var zyx = 1

func main() {

	fmt.Println("xyz=", xyz)
	fmt.Print("zyx=", zyx)
	ggg := 3
	fmt.Println("\nFirst ggg=", ggg)

	for ggg = 1; ggg < 20; ggg++ {

		fmt.Println(ggg)
	}

	fmt.Println("")
	foo2()
	foo()
}

func foo2() {

	fmt.Println("Foo says xyz=", xyz+10)
	fmt.Println("Foo says zyx=", zyx-5)

}
