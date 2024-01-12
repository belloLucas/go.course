package main

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
	a := "X"
	println(a, f)
}