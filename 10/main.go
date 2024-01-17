package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	lucas := Client{
		Name:   "Lucas",
		Age:    21,
		Active: true,
	}
	fmt.Printf("O nome do cliente é %s, a idade dele é %v. Está ativo? %v\n", lucas.Name, lucas.Age, lucas.Active)
}