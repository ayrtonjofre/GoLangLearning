package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	age, _ := reader.ReadString('\n')
	//age := 18

	/*
		// IF e ELSE case:

		if age <= 18 {
			fmt.Println("Yes, you can vote!")
		} else {
			fmt.Println("No, you can't vote!")
		}
	*/

	// SWITCH case:

	switch age {
	case "16":
		fmt.Println("Voce esta pronto para a faculdade")
	case "18":
		fmt.Println("Pare de correr atras de mulhe")
	case "20":
		fmt.Println("Vai arrumar um trab")
	default:
		fmt.Println("Ao menos voce esta vivo?")
	}
}
