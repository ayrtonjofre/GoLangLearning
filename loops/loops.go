package main

import "fmt"

func main() {
	//Forma normal de se fazer um loop como em qualquer outra linguagem
	/*
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}
	*/

	//Forma atual de se fazer onde tenho um controlador de saida (iteracao)

	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}
}
