package main

import "fmt"

func main() {
	fmt.Println("Hello class, let's hear something from foo")
	foo()
	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("Hello, I am in FOO")
}

func bar()  {
	fmt.Println("and then we exited...")
}
