package main

import "fmt"

func main() {
	fmt.Println(soma(10, 10))
}

func soma(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func times(a int, b int) int {
	return a * b
}

func div(a int, b int) int {
	return a / b
}
