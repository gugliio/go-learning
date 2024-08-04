package main

import "fmt"

func main() {
	fmt.Println("----------------------------------")
	fmt.Println("Strings")
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	fmt.Println("----------------------------------")
	fmt.Println("Boolean")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println("----------------------------------")
}
