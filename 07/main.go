package main

import "fmt"

func main() {
	salarios := map[string]int{"Lucas": 2200, "Fulano": 1500, "Ciclano": 3500}

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	  }

	// delete(salarios, "Ciclano")
	// salarios["Beltrano"] = 3200

	// sal := make(map[string]int)
	// sal["Fulano"] = 1412

	// sal1 := map[string]int{}
	// sal1["Lucas"] = 1500
}