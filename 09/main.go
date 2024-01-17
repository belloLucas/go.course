package main

import "fmt"

func main() {
	total := func() int {
		return sum(1,3,6,9,12,15,18,21,23, 29, 38, 42) * 2
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}