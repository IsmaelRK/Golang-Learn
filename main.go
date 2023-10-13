package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func main() {

	fmt.Println("HELLO WORLD!")

	var idade int
	idade = 16

	var nome string
	nome = "Ismael"

	nome2 := "Another Ismael"

	fmt.Println(idade, nome, nome2)

	if idade >= 18 {
		fmt.Println("Maior")
	} else {
		fmt.Println("menor")
	}

	fmt.Println(soma(5, 2))

	type Pessoas struct {
		Nome  string
		Idade int
	}
	
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	pessoa := Pessoas{Nome: "Mario", Idade: 25}
	fmt.Println(pessoa.Nome, "tem", pessoa.Idade, "anos")

}
