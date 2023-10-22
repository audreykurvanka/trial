package main

import "fmt"

func main() {
	fmt.Println(f(1.6))
	fmt.Println(f(-1.4))
	fmt.Println(f(0))
}

func f(x float32) float32 {
	var total float32
	total = ((x * x * x) + (3 * x)) / ((x * x * x * x) - (3 * x * x) + 4)
	return total
}
