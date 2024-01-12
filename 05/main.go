package main

import "fmt"

const text = "hello world"

type ID int

var (
	b bool    = true
	c int     = 24
	d string  = "Lucas"
	e float64 = 21.10
	f ID      = 1
)

func main() {
	var array [3]int
	array[0] = 10
	array[1] = 20
	array[2] = 30
	fmt.Println(array[0])

	for i, v := range array {
		fmt.Printf("O índice %d contém o valor %d\n", i, v)
	}
}