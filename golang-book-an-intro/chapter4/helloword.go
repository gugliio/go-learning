package main

import "fmt"

var x string = "Hello World"

func main() {
	fmt.Println(x)

	dogsName := "Max"
	fmt.Println("My dog's name is", dogsName)

	example()
}

func f() {
	fmt.Println(x)
}

func example() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
