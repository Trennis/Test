package main

import (
	"fmt"
)

var x = 10
var z int

func main() {
	type T struct {
		i    int
		f    float64
		next *T
	}
	for i := 0; i < 100; i++ {
		if i%5 == 0 {
			fmt.Println(i)

		}
	}
	y := 20
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println("x+y=", x+y)
	fmt.Println("z", z)
	fmt.Println("z+x+y=", z+x+y)
	var t T

	t.i = 1
	t.f = 2.5

	fmt.Println(t.i, t.f)
	fmt.Println(t)
	fmt.Println("*********************")
	fmt.Printf("%#v\n%+v\n%v", t, t, t)
	fmt.Println("*********************")
	fmt.Printf("%#v\n%+v\n%v\t%T", y, y, y, y)

}
